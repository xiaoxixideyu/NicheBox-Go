package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"nichebox/service/feed/api/internal/config"
	"nichebox/service/feed/rpc/feedclient"
	"nichebox/service/feed/rpc/pb/feed"
)

type ServiceContext struct {
	Config  config.Config
	FeedRpc feed.FeedClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		FeedRpc: feedclient.NewFeed(zrpc.MustNewClient(c.FeedRpc)),
	}
}
