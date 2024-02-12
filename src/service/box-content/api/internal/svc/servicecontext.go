package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"nichebox/service/box-content/api/internal/config"
	"nichebox/service/box-content/rpc/boxcontent"
	box_content "nichebox/service/box-content/rpc/pb/box-content"
)

type ServiceContext struct {
	Config        config.Config
	BoxContentRpc box_content.BoxContentClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		BoxContentRpc: boxcontent.NewBoxContent(zrpc.MustNewClient(c.BoxContentRpc)),
	}
}
