package redis

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"nichebox/service/email/model"
	"time"
)

type RedisInterface struct {
	rds *redis.Redis
}

func NewRedisInterface(hosts []string, deployType, pass string, tls, nonBlock bool, pingTimeout int) (model.EmailRedisInterface, error) {
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

func (r *RedisInterface) SetVerificationCodeRegister(ctx context.Context, destination, code string, expiration int) error {
	key := KeyPrefixUser + KeyRegisterCode + destination
	return r.rds.SetexCtx(ctx, key, code, expiration)
}

func (r *RedisInterface) SetVerificationCodePWD(ctx context.Context, destination, code string, expiration int) error {
	key := KeyPrefixUser + KeyPWDCode + destination
	return r.rds.SetexCtx(ctx, key, code, expiration)
}
