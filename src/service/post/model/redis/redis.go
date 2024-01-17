package redis

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"nichebox/service/post/model"
	"strconv"
	"time"
)

type RedisInterface struct {
	rds *redis.Redis
}

func NewRedisInterface(hosts []string, deployType, pass string, tls, nonBlock bool, pingTimeout int) (model.PostCacheInterface, error) {
	conf := redis.RedisConf{
		// todo: judge node or cluster
		Host:        hosts[0],
		Type:        deployType,
		Pass:        pass,
		Tls:         tls,
		NonBlock:    nonBlock,
		PingTimeout: time.Duration(pingTimeout) * time.Second,
	}
	r := &RedisInterface{
		rds: redis.MustNewRedis(conf),
	}
	return r, nil
}

func (r *RedisInterface) GetUserView(ctx context.Context, postID int64) (int64, error) {
	key := KeyPrefixPost + KeyUserView + strconv.FormatInt(postID, 10)
	val, err := r.rds.PfcountCtx(ctx, key)
	return val, err
}

func (r *RedisInterface) IncrUserView(ctx context.Context, postID int64, uid int64) error {
	key := KeyPrefixPost + KeyUserView + strconv.FormatInt(postID, 10)
	_, err := r.rds.PfaddCtx(ctx, key, uid)
	return err
}
