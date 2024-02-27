package svc

import (
	"log"
	"nichebox/common/leaf"
	"nichebox/service/box_info/model"
	"nichebox/service/box_info/model/mysql"
	"nichebox/service/box_info/rpc/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	BoxInterface model.BoxInterface
	Leaves       *leaf.Leaves
}

func NewServiceContext(c config.Config) *ServiceContext {
	boxInterface, err := mysql.NewMysqlInterface(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("failed to create user interface, err: %v\n", err)
		return nil
	}
	leaves, err := leaf.NewLeavesMysql(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("failed to create leaves, err: %v\n", err)
		return nil
	}

	return &ServiceContext{
		Config:       c,
		BoxInterface: boxInterface,
		Leaves:       leaves,
	}
}
