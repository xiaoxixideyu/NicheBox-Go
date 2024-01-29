package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"nichebox/service/post/api/internal/config"
	"nichebox/service/post/rpc/postclient"
	"nichebox/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
	PostRpc postclient.Post
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		PostRpc: postclient.NewPost(zrpc.MustNewClient(c.PostRpc)),
	}
}
