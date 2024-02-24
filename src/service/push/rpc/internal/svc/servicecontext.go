package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"nichebox/service/push/rpc/internal/config"
)

type ServiceContext struct {
	Config                   config.Config
	KqPushToUserPusherClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	kqPushToUserPusher := kq.NewPusher(c.KqPushToUserPusherConf.Brokers, c.KqPushToUserPusherConf.Topic)
	return &ServiceContext{
		Config:                   c,
		KqPushToUserPusherClient: kqPushToUserPusher,
	}
}
