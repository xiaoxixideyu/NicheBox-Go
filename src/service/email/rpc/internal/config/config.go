package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	ServerMail struct {
		Address  string
		Password string
		Host     string
		Port     int
	}
}
