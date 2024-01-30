package logic

import (
	"context"

	"nichebox/service/comment/rpc/internal/svc"
	"nichebox/service/comment/rpc/pb/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsFromSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentsFromSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsFromSubjectLogic {
	return &GetCommentsFromSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentsFromSubjectLogic) GetCommentsFromSubject(in *comment.GetCommentsFromSubjectRequest) (*comment.GetCommentsFromSubjectResponse, error) {
	// todo: add your logic here and delete this line

	return &comment.GetCommentsFromSubjectResponse{}, nil
}
