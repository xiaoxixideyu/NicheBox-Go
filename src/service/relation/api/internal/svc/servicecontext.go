package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"nichebox/service/relation/api/internal/config"
	"nichebox/service/relation/rpc/pb/relation"
	"nichebox/service/relation/rpc/relationclient"
)

type ServiceContext struct {
	Config      config.Config
	RelationRpc relation.RelationClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RelationRpc: relationclient.NewRelation(zrpc.MustNewClient(c.RelationRpc)),
	}
}
