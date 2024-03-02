package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"log"
	"nichebox/service/relation/model"
	"nichebox/service/relation/model/mysql"
	"nichebox/service/relation/model/redis"
	"nichebox/service/relation/rpc/internal/config"
)

type ServiceContext struct {
	Config                       config.Config
	RelationInterface            model.RelationInterface
	RelationCacheInterface       model.RelationCacheInterface
	KqRebuildRelationCachePusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	relationInterface, err := mysql.NewMysqlInterface(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("failed to create relation interface, err: %v\n", err)
		return nil
	}
	relationRedisInterface, err := redis.NewRedisInterface(c.CacheRedis.Host, c.CacheRedis.Type, c.CacheRedis.Pass, c.CacheRedis.Tls, c.CacheRedis.NonBlock, c.CacheRedis.PingTimeout, c.CacheRedis.BloomFilterBits)
	if err != nil {
		log.Printf("failed to create relation redis interface, err:%v\n", err)
		return nil
	}
	kqRebuildRelationCachePusher := kq.NewPusher(c.KqRebuildRelationCachePusherConf.Brokers, c.KqRebuildRelationCachePusherConf.Topic)
	return &ServiceContext{
		Config:                       c,
		RelationInterface:            relationInterface,
		RelationCacheInterface:       relationRedisInterface,
		KqRebuildRelationCachePusher: kqRebuildRelationCachePusher,
	}
}
