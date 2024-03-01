package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"nichebox/service/box_user/rpc/internal/config"
	"nichebox/service/box_user/rpc/internal/mqs"
	"nichebox/service/box_user/rpc/internal/server"
	"nichebox/service/box_user/rpc/internal/svc"
	"nichebox/service/box_user/rpc/pb/boxuser"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/boxuser.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	// mq
	go func() {
		defer func() {
			if p := recover(); p != nil {
				log.Printf("mq panic:%v", p)
			}
		}()

		serviceGroup := service.NewServiceGroup()
		defer serviceGroup.Stop()

		for _, mq := range mqs.Consumers(c, context.Background(), ctx) {
			serviceGroup.Add(mq)
		}
		serviceGroup.Start()

	}()

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		boxuser.RegisterBoxUserServer(grpcServer, server.NewBoxUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
