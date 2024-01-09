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
	CacheRedis struct {
		Host        []string
		Type        string
		Pass        string
		Tls         bool
		NonBlock    bool
		PingTimeout int
	}
}
