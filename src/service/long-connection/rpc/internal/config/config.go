package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	ServerConf struct {
		TCPPort      int
		PingInterval int
	}
	KqPushToUserConsumerConf   kq.KqConf
	KqPushToDeviceConsumerConf kq.KqConf
}
