package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"log"
	"nichebox/common/snowflake"
	"nichebox/service/post/model"
	"nichebox/service/post/model/mysql"
	"nichebox/service/post/model/redis"
	"nichebox/service/post/rpc/internal/config"
)

type ServiceContext struct {
	Config                       config.Config
	PostInterface                model.PostInterface
	PostCacheInterface           model.PostCacheInterface
	KqUpdateUserViewPusherClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	postInterface, err := mysql.NewMysqlInterface(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("failed to create user interface, err: %v\n", err)
		return nil
	}
	postRedisInterface, err := redis.NewRedisInterface(c.CacheRedis.Host, c.CacheRedis.Type, c.CacheRedis.Pass, c.CacheRedis.Tls, c.CacheRedis.NonBlock, c.CacheRedis.PingTimeout)
	if err != nil {
		log.Printf("failed to create email redis interface, err:%v\n", err)
		return nil
	}
	err = snowflake.Init(c.Snowflake.MachineID)
	if err != nil {
		log.Printf("failed to initialize snowflake, err:%v\n", err)
		return nil
	}
	kqUpdateUserViewPusher := kq.NewPusher(c.KqUpdateUserViewPusherConf.Brokers, c.KqUpdateUserViewPusherConf.Topic)
	return &ServiceContext{
		Config:                       c,
		PostInterface:                postInterface,
		PostCacheInterface:           postRedisInterface,
		KqUpdateUserViewPusherClient: kqUpdateUserViewPusher,
	}
}
