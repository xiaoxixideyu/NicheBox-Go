package redis

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"nichebox/service/task/model"
	"strconv"

	"time"
)

type RedisInterface struct {
	rds *redis.Redis
}

func NewRedisInterface(hosts []string, deployType, pass string, tls, nonBlock bool, pingTimeout int) (model.TaskCacheInterface, error) {
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
	}
	return r, nil
}

func (r *RedisInterface) GetUserView(ctx context.Context, postID int64) (int64, error) {
	key := KeyPrefixPost + KeyUserView + strconv.FormatInt(postID, 10)
	val, err := r.rds.PfcountCtx(ctx, key)
	return val, err
}
