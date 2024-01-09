package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	ServerMail struct {
		Address  string
		Password string
		Host     string
		Port     int
	}
	EmailRpc zrpc.RpcClientConf
}
