package logic

import (
	"context"

	"nichebox/service/comment/api/internal/svc"
	"nichebox/service/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSubCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubCommentsLogic {
	return &GetSubCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubCommentsLogic) GetSubComments(req *types.GetSubCommentsRequest) (resp *types.GetSubCommentsResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
