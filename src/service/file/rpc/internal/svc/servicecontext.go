package svc

import (
	"log"
	"nichebox/common/snowflake"
	"nichebox/service/file/model"
	"nichebox/service/file/model/mysql"
	"nichebox/service/file/rpc/internal/config"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type ServiceContext struct {
	Config             config.Config
	OssClient          *oss.Client
	ImageFileInterface model.ImageFileInterface
}

func NewServiceContext(c config.Config) *ServiceContext {
	err := snowflake.Init(c.Snowflake.MachineID)
	if err != nil {
		log.Printf("failed to initialize snowflake, err:%v\n", err)
		os.Exit(1)
	}

	ossClient, err := oss.New(c.AliyunOSS.Endpoint, c.AliyunOSS.AccessKey, c.AliyunOSS.AccessKeySecret)
	if err != nil {
		log.Printf("Create Aliyun OSS failed, error: %v\n", err)
		os.Exit(1)
	}

	imageFileInterface, err := mysql.NewMysqlInterface(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("failed to create user interface, err: %v\n", err)
		return nil
	}

	return &ServiceContext{
		Config:             c,
		OssClient:          ossClient,
		ImageFileInterface: imageFileInterface,
	}
}
