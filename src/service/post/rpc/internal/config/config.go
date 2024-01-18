package config

import "github.com/zeromicro/go-zero/zrpc"

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
	Snowflake struct {
		MachineID int64
	}
	KqUpdateUserViewPusherConf struct {
		Brokers []string
		Topic   string
	}
	CacheRedis struct {
		Host            []string
		Type            string
		Pass            string
		Tls             bool
		NonBlock        bool
		PingTimeout     int
		BloomFilterBits uint
	}
	Salt string
}
