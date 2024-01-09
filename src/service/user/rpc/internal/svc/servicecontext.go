package svc

import (
	"log"
	"nichebox/common/snowflake"
	"nichebox/service/user/model"
	"nichebox/service/user/model/mysql"
	"nichebox/service/user/model/redis"
	"nichebox/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config             config.Config
	UserInterface      model.UserInterface
	UserRedisInterface model.UserRedisInterface
}

func NewServiceContext(c config.Config) *ServiceContext {
	userInterface, err := mysql.NewMysqlInterface(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("failed to create user interface, err: %v\n", err)
		return nil
	}
	userRedisInterface, err := redis.NewRedisInterface(c.CacheRedis.Host, c.CacheRedis.Type, c.CacheRedis.Pass, c.CacheRedis.Tls, c.CacheRedis.NonBlock, c.CacheRedis.PingTimeout)
	if err != nil {
		log.Printf("failed to create email redis interface, err:%v\n", err)
		return nil
	}
	err = snowflake.Init(c.Snowflake.MachineID)
	if err != nil {
		log.Printf("failed to initialize snowflake, err:%v\n", err)
		return nil
	}

	return &ServiceContext{
		Config:             c,
		UserInterface:      userInterface,
		UserRedisInterface: userRedisInterface,
	}
}
