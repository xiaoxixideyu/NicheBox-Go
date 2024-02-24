package mqs

import (
	"context"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"nichebox/service/long-connection/rpc/internal/config"
	"nichebox/service/long-connection/rpc/internal/svc"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.KqPushToUserConsumerConf, NewPushToUser(ctx, svcContext)),
	}

}
