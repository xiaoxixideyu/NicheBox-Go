package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"nichebox/service/email/api/internal/config"
	"nichebox/service/email/rpc/emailclient"
)

type ServiceContext struct {
	Config config.Config

	EmailRpc emailclient.Email
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		EmailRpc: emailclient.NewEmail(zrpc.MustNewClient(c.EmailRpc)),
	}
}
