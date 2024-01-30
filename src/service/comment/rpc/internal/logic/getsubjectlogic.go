package logic

import (
	"context"

	"nichebox/service/comment/rpc/internal/svc"
	"nichebox/service/comment/rpc/pb/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubjectLogic {
	return &GetSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSubjectLogic) GetSubject(in *comment.GetSubjectRequest) (*comment.GetSubjectResponse, error) {
	// todo: add your logic here and delete this line

	return &comment.GetSubjectResponse{}, nil
}
