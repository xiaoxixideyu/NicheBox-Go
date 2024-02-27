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
	AliyunOSS struct {
		BucketName      string
		Endpoint        string
		AccessKey       string
		AccessKeySecret string
	}
	Snowflake struct {
		MachineID int64
	}
	Image struct {
		Quality float32
	}
}
