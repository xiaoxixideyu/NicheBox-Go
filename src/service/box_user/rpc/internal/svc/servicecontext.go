package svc

import (
	"log"
	"nichebox/service/box_user/model"
	"nichebox/service/box_user/model/mysql"
	"nichebox/service/box_user/model/redis"
	"nichebox/service/box_user/rpc/internal/config"

	"github.com/zeromicro/go-queue/kq"
)

type ServiceContext struct {
	Config                           config.Config
	BoxUserInterface                 model.BoxUserInterface
	BoxUserCaCheInterface            model.BoxUserCacheInterface
	KqRemoveCacheBoxUserPusherClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	boxUserInterface, err := mysql.NewMysqlInterface(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("failed to create user interface, err: %v\n", err)
		return nil
	}
	boxUserCacheInterface, err := redis.NewRedisInterface(c.CacheRedis.Host, c.CacheRedis.Type, c.CacheRedis.Pass, c.CacheRedis.Tls, c.CacheRedis.NonBlock, c.CacheRedis.PingTimeout)
	if err != nil {
		log.Printf("failed to create email redis interface, err:%v\n", err)
		return nil
	}

	return &ServiceContext{
		Config:                           c,
		BoxUserInterface:                 boxUserInterface,
		BoxUserCaCheInterface:            boxUserCacheInterface,
		KqRemoveCacheBoxUserPusherClient: kq.NewPusher(c.KqRemoveCacheBoxUserPusherConf.Brokers, c.KqRemoveCacheBoxUserPusherConf.Topic),
	}
}
