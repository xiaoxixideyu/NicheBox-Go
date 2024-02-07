package mqs

import (
	"context"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"nichebox/service/comment/rpc/internal/config"
	"nichebox/service/comment/rpc/internal/svc"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.KqRebuildCacheSubjectCommentIndexConsumerConf, NewRebuildCacheSubjectCommentIndex(ctx, svcContext)),
		kq.MustNewQueue(c.KqRebuildCacheInnerFloorCommentIndexConsumerConf, NewRebuildCacheInnerFloorCommentIndex(ctx, svcContext)),
	}

}
