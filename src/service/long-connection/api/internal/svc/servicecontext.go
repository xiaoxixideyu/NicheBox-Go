package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"nichebox/service/long-connection/api/internal/config"
	"nichebox/service/long-connection/api/internal/middleware"
	"nichebox/service/long-connection/rpc/longconnclient"
	"nichebox/service/long-connection/rpc/pb/longConn"
)

type ServiceContext struct {
	Config                  config.Config
	LongConnRpc             longConn.LongConnClient
	RemoteAddressMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                  c,
		LongConnRpc:             longconnclient.NewLongConn(zrpc.MustNewClient(c.LongConnRpc)),
		RemoteAddressMiddleware: middleware.NewRemoteAddressMiddleware().Handle,
	}
}
