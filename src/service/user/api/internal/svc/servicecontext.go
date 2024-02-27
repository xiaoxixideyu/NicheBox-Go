package svc

import (
	"nichebox/service/email/rpc/emailclient"
	"nichebox/service/file/rpc/fileclient"
	"nichebox/service/user/api/internal/config"
	"nichebox/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	UserRpc  userclient.User
	EmailRpc emailclient.Email
	FileRpc  fileclient.File
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		UserRpc:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		EmailRpc: emailclient.NewEmail(zrpc.MustNewClient(c.EmailRpc)),
		FileRpc:  fileclient.NewFile(zrpc.MustNewClient(c.FileRpc)),
	}
}
