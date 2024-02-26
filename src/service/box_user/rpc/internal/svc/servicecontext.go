package svc

import (
	"log"
	"nichebox/service/box_user/model"
	"nichebox/service/box_user/model/mysql"
	"nichebox/service/box_user/rpc/internal/config"
)

type ServiceContext struct {
	Config           config.Config
	BoxUserInterface model.BoxUserInterface
}

func NewServiceContext(c config.Config) *ServiceContext {
	boxUserInterface, err := mysql.NewMysqlInterface(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("failed to create user interface, err: %v\n", err)
		return nil
	}
	return &ServiceContext{
		Config:           c,
		BoxUserInterface: boxUserInterface,
	}
}
