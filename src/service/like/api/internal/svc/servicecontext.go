package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"nichebox/service/like/api/internal/config"
	"nichebox/service/like/rpc/likeclient"
	"nichebox/service/like/rpc/pb/like"
)

type ServiceContext struct {
	Config  config.Config
	LikeRpc like.LikeClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		LikeRpc: likeclient.NewLike(zrpc.MustNewClient(c.LikeRpc)),
	}
}
