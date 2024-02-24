package svc

import (
	"context"
	"nichebox/service/long-connection/rpc/internal/access/session"
	"nichebox/service/long-connection/rpc/internal/config"
)

type ServiceContext struct {
	Config         config.Config
	SessionManager session.Manager
}

func NewServiceContext(c config.Config) *ServiceContext {
	m := session.NewSessionManager(context.Background(), c.ServerConf.PingInterval)
	err := m.Start(c)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:         c,
		SessionManager: m,
	}
}
