package logic

import (
	"context"

	"nichebox/service/post/rpc/internal/svc"
	"nichebox/service/post/rpc/pb/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseUserViewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncreaseUserViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseUserViewLogic {
	return &IncreaseUserViewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IncreaseUserViewLogic) IncreaseUserView(in *post.IncreaseUserViewRequest) (*post.IncreaseUserViewResponse, error) {
	// todo: 如果是游客的话，调用此RPC的visitorID应该换成ip地址
	err := l.svcCtx.PostCacheInterface.IncrUserView(l.ctx, in.PostID, in.VisitorID)
	if err != nil {
		return nil, err
	}
	// todo: 新增定时任务

	return &post.IncreaseUserViewResponse{}, nil
}
