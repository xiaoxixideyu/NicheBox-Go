package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nichebox/common/biz"
	"strconv"

	"nichebox/service/box-content/rpc/internal/svc"
	"nichebox/service/box-content/rpc/pb/box-content"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostListLogic {
	return &GetPostListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostListLogic) GetPostList(in *box_content.GetPostListRequest) (*box_content.GetPostListResponse, error) {
	kvs, err := l.svcCtx.BoxContentCacheInterface.GetPostIDsCtx(l.ctx, in.BoxID, int(in.Page), int(in.Size), in.Order)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			// todo: 请求盒子信息服务查看内容数量，如果是0说明盒子无信息，然后返回，否则返回错误
			return &box_content.GetPostListResponse{IDs: nil}, nil
		} else if errors.Is(err, biz.ErrRedisOutOfBounds) {
			return nil, status.Error(codes.OutOfRange, err.Error())
		} else if errors.Is(err, biz.ErrRedisUnknownOrder) {
			return nil, err
		}

		l.Logger.Errorf("[Redis] Get post ids error", err)
		// todo: 找个方法查数据库

	}

	ids := make([]int64, 0, len(kvs))

	for _, v := range kvs {
		id, err := strconv.ParseInt(v.Key, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}

	return &box_content.GetPostListResponse{IDs: ids}, nil
}
