package redis

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"nichebox/service/user/model"
	"time"
)

type RedisInterface struct {
	rds *redis.Redis
}

func NewRedisInterface(hosts []string, deployType, pass string, tls, nonBlock bool, pingTimeout int) (model.UserRedisInterface, error) {
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

func (r RedisInterface) GetVerificationCode(ctx context.Context, key string) (string, error) {
	return r.rds.GetCtx(ctx, key)
}
