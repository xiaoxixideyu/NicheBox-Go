package redis

import (
	"context"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"nichebox/service/post/model"
	"strconv"
	"strings"
	"sync"
	"time"
)

type RedisInterface struct {
	rds *redis.Redis
	*filter
}

type filter struct {
	f     *bloom.Filter
	today int
	bits  uint

	mu sync.RWMutex
}

func NewRedisInterface(hosts []string, deployType, pass string, tls, nonBlock bool, pingTimeout int, bloomFilterBits uint) (model.PostCacheInterface, error) {
	conf := redis.RedisConf{
		// todo: judge node or cluster
		Host:        hosts[0],
		Type:        deployType,
		Pass:        pass,
		Tls:         tls,
		NonBlock:    nonBlock,
		PingTimeout: time.Duration(pingTimeout) * time.Second,
	}
	rds := redis.MustNewRedis(conf)
	r := &RedisInterface{
		rds: rds,
		filter: &filter{
			bits: bloomFilterBits,
		},
	}
	return r, nil
}

func (r *RedisInterface) GetUserView(ctx context.Context, postID int64) (int64, error) {
	key := KeyPrefixPost + KeyUserView + strconv.FormatInt(postID, 10)
	val, err := r.rds.PfcountCtx(ctx, key)
	return val, err
}

func (r *RedisInterface) IncrUserView(ctx context.Context, postID int64, visitorID int64) error {
	key := KeyPrefixPost + KeyUserView + strconv.FormatInt(postID, 10)
	_, err := r.rds.PfaddCtx(ctx, key, visitorID)
	return err
}

func (r *RedisInterface) BloomCheckPostExists(ctx context.Context, postID int64) (bool, error) {
	bf := r.bloomFilter()

	exists, err := bf.ExistsCtx(ctx, []byte(strconv.FormatInt(postID, 10)))
	if err != nil {
		return true, err
	}
	return exists, nil
}

func (r *RedisInterface) BloomAddPost(ctx context.Context, postID int64) error {
	bf := r.bloomFilter()

	err := bf.AddCtx(ctx, []byte(strconv.FormatInt(postID, 10)))
	return err
}

func (r *RedisInterface) bloomFilter() *bloom.Filter {
	r.filter.mu.RLock()

	// 未初始化或者filter时间已跨天
	if time.Now().Day() != r.filter.today {
		r.filter.mu.RUnlock()
		// 加写锁
		r.filter.mu.Lock()
		defer r.filter.mu.Unlock()

		// 单例的double check
		if time.Now().Day() == r.filter.today {
			return r.filter.f
		}

		now := time.Now()
		date := strings.Trim(now.Format("2006-01-02"), "-")
		key := KeyPrefixPost + KeyPrefixBloomFilter + date

		r.filter.f = bloom.New(r.rds, key, r.filter.bits)
		r.filter.today = now.Day()

		r.rds.Expire(key, UpdateUserViewBloomFilterExpiration)
	} else {
		r.filter.mu.RUnlock()
	}

	return r.filter.f
}
