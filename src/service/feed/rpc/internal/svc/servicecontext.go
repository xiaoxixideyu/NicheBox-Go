package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
	"log"
	"nichebox/common/snowflake"
	"nichebox/service/feed/model"
	"nichebox/service/feed/model/mysql"
	"nichebox/service/feed/rpc/internal/config"
	"nichebox/service/relation/rpc/pb/relation"
	"nichebox/service/relation/rpc/relationclient"
)

type ServiceContext struct {
	Config                            config.Config
	FeedInterface                     model.FeedInterface
	KqDeliverFeedToOutboxPusherClient *kq.Pusher
	RelationRpc                       relation.RelationClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	feedInterface, err := mysql.NewMysqlInterface(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("failed to create feed interface, err: %v\n", err)
		return nil
	}
	err = snowflake.Init(c.Snowflake.MachineID)
	if err != nil {
		log.Printf("failed to initialize snowflake, err:%v\n", err)
		return nil
	}
	kqDeliverFeedToOutboxPusherClient := kq.NewPusher(c.KqDeliverFeedToOutboxPusherConf.Brokers, c.KqDeliverFeedToOutboxPusherConf.Topic)

	return &ServiceContext{
		Config:                            c,
		FeedInterface:                     feedInterface,
		KqDeliverFeedToOutboxPusherClient: kqDeliverFeedToOutboxPusherClient,
		RelationRpc:                       relationclient.NewRelation(zrpc.MustNewClient(c.RelationRpc)),
	}
}
