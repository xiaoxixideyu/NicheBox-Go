package svc

import (
	"log"
	"nichebox/service/email/model"
	"nichebox/service/email/model/redis"
	"nichebox/service/email/rpc/internal/config"
)

type ServiceContext struct {
	Config              config.Config
	EmailRedisInterface model.EmailRedisInterface
}

func NewServiceContext(c config.Config) *ServiceContext {
	emailRedisInterface, err := redis.NewRedisInterface(c.CacheRedis.Host, c.CacheRedis.Type, c.CacheRedis.Pass, c.CacheRedis.Tls, c.CacheRedis.NonBlock, c.CacheRedis.PingTimeout)
	if err != nil {
		log.Printf("failed to create email redis interface, err:%v\n", err)
		return nil
	}
	return &ServiceContext{
		Config:              c,
		EmailRedisInterface: emailRedisInterface,
	}
}
