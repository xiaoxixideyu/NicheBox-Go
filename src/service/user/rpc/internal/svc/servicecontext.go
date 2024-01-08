package svc

import (
	"log"
	"nichebox/common/snowflake"
	"nichebox/service/user/model"
	"nichebox/service/user/model/mysql"
	"nichebox/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	UserInterface model.UserInterface
}

func NewServiceContext(c config.Config) *ServiceContext {
	userInterface, err := mysql.NewMysqlInterface(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("failed to create user interface, err: %v\n", err)
		return nil
	}
	err = snowflake.Init(c.Snowflake.MachineID)
	if err != nil {
		log.Printf("failed to initialize snowflake, err:%v\n", err)
		return nil
	}

	return &ServiceContext{
		Config:        c,
		UserInterface: userInterface,
	}
}
