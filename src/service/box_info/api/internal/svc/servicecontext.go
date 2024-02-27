package svc

import (
	"nichebox/service/box_info/api/internal/config"
	"nichebox/service/box_info/rpc/boxinfoclient"
	"nichebox/service/box_user/rpc/boxuserclient"
	"nichebox/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	UserRpc    userclient.User
	BoxInfoRpc boxinfoclient.BoxInfo
	BoxUserRpc boxuserclient.BoxUser
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		BoxInfoRpc: boxinfoclient.NewBoxInfo(zrpc.MustNewClient(c.BoxInfoRpc)),
		BoxUserRpc: boxuserclient.NewBoxUser(zrpc.MustNewClient(c.BoxUserRpc)),
	}
}
