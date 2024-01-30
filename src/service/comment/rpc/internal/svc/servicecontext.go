package svc

import (
	"log"
	"nichebox/service/comment/model"
	"nichebox/service/comment/model/mysql"
	"nichebox/service/comment/model/redis"
	"nichebox/service/comment/rpc/internal/config"
)

type ServiceContext struct {
	Config                config.Config
	CommentInterface      model.CommentInterface
	CommentCacheInterface model.CommentCacheInterface
}

func NewServiceContext(c config.Config) *ServiceContext {
	commentInterface, err := mysql.NewMysqlInterface(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("failed to create like interface, err: %v\n", err)
		return nil
	}
	commentRedisInterface, err := redis.NewRedisInterface(c.CacheRedis.Host, c.CacheRedis.Type, c.CacheRedis.Pass, c.CacheRedis.Tls, c.CacheRedis.NonBlock, c.CacheRedis.PingTimeout)
	if err != nil {
		log.Printf("failed to create like redis interface, err:%v\n", err)
		return nil
	}
	return &ServiceContext{
		Config:                c,
		CommentInterface:      commentInterface,
		CommentCacheInterface: commentRedisInterface,
	}
}
