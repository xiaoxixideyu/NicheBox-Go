package logic

import (
	"context"
	"nichebox/service/post/rpc/pb/post"

	"nichebox/service/post/api/internal/svc"
	"nichebox/service/post/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseUserViewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIncreaseUserViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseUserViewLogic {
	return &IncreaseUserViewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IncreaseUserViewLogic) IncreaseUserView(req *types.IncreaseUserViewRequest) (resp *types.IncreaseUserViewResponse, err error) {
	in := post.IncreaseUserViewRequest{
		PostID:    req.PostID,
		VisitorID: req.Uid,
	}

	_, _ = l.svcCtx.PostRpc.IncreaseUserView(l.ctx, &in)

	return
}
