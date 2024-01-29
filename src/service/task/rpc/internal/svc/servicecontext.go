package svc

import (
	"github.com/robfig/cron/v3"
	"log"
	"nichebox/service/task/model"
	"nichebox/service/task/model/mysql"
	"nichebox/service/task/model/redis"
	"nichebox/service/task/rpc/internal/config"
	"sync"
)

type ServiceContext struct {
	Config             config.Config
	UpdateUserViewCond *sync.Cond
	Cron               *cron.Cron
	TaskInterface      model.TaskInterface
	TaskCacheInterface model.TaskCacheInterface
}

func NewServiceContext(c config.Config) *ServiceContext {
	taskInterface, err := mysql.NewMysqlInterface(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("failed to create task interface, err: %v\n", err)
		return nil
	}
	taskRedisInterface, err := redis.NewRedisInterface(c.CacheRedis.Host, c.CacheRedis.Type, c.CacheRedis.Pass, c.CacheRedis.Tls, c.CacheRedis.NonBlock, c.CacheRedis.PingTimeout)
	if err != nil {
		log.Printf("failed to create task redis interface, err:%v\n", err)
		return nil
	}
	rootCron := cron.New()
	rootCron.Start()
	return &ServiceContext{
		Config:             c,
		UpdateUserViewCond: sync.NewCond(&sync.Mutex{}),
		Cron:               rootCron,
		TaskInterface:      taskInterface,
		TaskCacheInterface: taskRedisInterface,
	}
}
