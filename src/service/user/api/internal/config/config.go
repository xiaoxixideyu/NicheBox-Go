package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	Auth struct {
		AccessSecret  string
		AccessExpire  int64
		RefreshExpire int64
	}

	UserRpc  zrpc.RpcClientConf
	EmailRpc zrpc.RpcClientConf
	FileRpc  zrpc.RpcClientConf

	File struct {
		MaxMemory int64
	}
}
