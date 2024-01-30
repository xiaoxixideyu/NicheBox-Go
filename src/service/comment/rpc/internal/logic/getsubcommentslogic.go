package logic

import (
	"context"

	"nichebox/service/comment/rpc/internal/svc"
	"nichebox/service/comment/rpc/pb/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSubCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubCommentsLogic {
	return &GetSubCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSubCommentsLogic) GetSubComments(in *comment.GetSubCommentsRequest) (*comment.GetSubCommentsResponse, error) {
	// todo: add your logic here and delete this line

	return &comment.GetSubCommentsResponse{}, nil
}
