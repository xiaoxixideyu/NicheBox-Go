package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"nichebox/service/task/rpc/internal/config"
	"nichebox/service/task/rpc/internal/cronx"
	"nichebox/service/task/rpc/internal/mqs"
	"nichebox/service/task/rpc/internal/server"
	"nichebox/service/task/rpc/internal/svc"
	"nichebox/service/task/rpc/pb/task"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/task.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		task.RegisterTaskServer(grpcServer, server.NewTaskServer(svcCtx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// mq
	go func() {
		defer func() {
			if p := recover(); p != nil {
				log.Printf("mq panic:%v", p)
			}
		}()

		serviceGroup := service.NewServiceGroup()
		defer serviceGroup.Stop()

		for _, mq := range mqs.Consumers(c, context.Background(), svcCtx) {
			serviceGroup.Add(mq)
		}
		serviceGroup.Start()

	}()

	// cron
	userViewTask := cronx.NewUpdateUserView(context.Background(), svcCtx)
	userViewTask.AddUpdateUserViewTask()
	updateBoxContentTask := cronx.NewUpdateBoxContent(context.Background(), svcCtx)
	updateBoxContentTask.AddUpdateBoxContentTask()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()

}
