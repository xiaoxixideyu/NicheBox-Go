package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	KqPushToUserPusherConf struct {
		Brokers []string
		Topic   string
	}
}
