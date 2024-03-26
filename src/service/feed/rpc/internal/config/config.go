package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataBase        string
		Username        string
		Password        string
		Host            string
		Port            string
		MaxIdleConns    int
		MaxOpenConns    int
		ConnMaxLifeTime int
	}
	CacheRedis struct {
		Host        []string
		Type        string
		Pass        string
		Tls         bool
		NonBlock    bool
		PingTimeout int
	}
	Snowflake struct {
		MachineID int64
	}
	RelationRpc                     zrpc.RpcClientConf
	KqDeliverFeedToOutboxPusherConf struct {
		Brokers []string
		Topic   string
	}
	KqDeliverFeedToOutboxConsumerConf kq.KqConf
}
