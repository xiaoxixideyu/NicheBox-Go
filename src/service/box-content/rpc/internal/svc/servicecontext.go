package svc

import (
	"log"
	"nichebox/service/box-content/model"
	"nichebox/service/box-content/model/redis"
	"nichebox/service/box-content/rpc/internal/config"
)

type ServiceContext struct {
	Config                   config.Config
	BoxContentCacheInterface model.BoxContentCacheInterface
}

func NewServiceContext(c config.Config) *ServiceContext {
	boxContentRedisInterface, err := redis.NewRedisInterface(c.CacheRedis.Host, c.CacheRedis.Type, c.CacheRedis.Pass, c.CacheRedis.Tls, c.CacheRedis.NonBlock, c.CacheRedis.PingTimeout)
	if err != nil {
		log.Printf("failed to create box-content redis interface, err:%v\n", err)
		return nil
	}
	return &ServiceContext{
		Config:                   c,
		BoxContentCacheInterface: boxContentRedisInterface,
	}
}
