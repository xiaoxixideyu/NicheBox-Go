package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"log"
	"nichebox/service/like/model"
	"nichebox/service/like/model/mysql"
	"nichebox/service/like/model/redis"
	"nichebox/service/like/rpc/internal/config"
)

type ServiceContext struct {
	Config                         config.Config
	LikeInterface                  model.LikeInterface
	LikeCacheInterface             model.LikeCacheInterface
	KqUpdateCommentLikeCountPusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	likeInterface, err := mysql.NewMysqlInterface(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("failed to create like interface, err: %v\n", err)
		return nil
	}
	likeRedisInterface, err := redis.NewRedisInterface(c.CacheRedis.Host, c.CacheRedis.Type, c.CacheRedis.Pass, c.CacheRedis.Tls, c.CacheRedis.NonBlock, c.CacheRedis.PingTimeout)
	if err != nil {
		log.Printf("failed to create like redis interface, err:%v\n", err)
		return nil
	}
	kqUpdateCommentLikeCountPusher := kq.NewPusher(c.KqUpdateCommentLikeCountPusherConf.Brokers, c.KqUpdateCommentLikeCountPusherConf.Topic)
	return &ServiceContext{
		Config:                         c,
		LikeInterface:                  likeInterface,
		LikeCacheInterface:             likeRedisInterface,
		KqUpdateCommentLikeCountPusher: kqUpdateCommentLikeCountPusher,
	}
}
