package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"nichebox/service/feed/rpc/internal/mqs"

	"nichebox/service/feed/rpc/internal/config"
	"nichebox/service/feed/rpc/internal/server"
	"nichebox/service/feed/rpc/internal/svc"
	"nichebox/service/feed/rpc/pb/feed"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/feed.yaml", "the config file")

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
		feed.RegisterFeedServer(grpcServer, server.NewFeedServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
