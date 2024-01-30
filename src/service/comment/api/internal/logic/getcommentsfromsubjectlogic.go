package logic

import (
	"context"

	"nichebox/service/comment/api/internal/svc"
	"nichebox/service/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsFromSubjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentsFromSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsFromSubjectLogic {
	return &GetCommentsFromSubjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentsFromSubjectLogic) GetCommentsFromSubject(req *types.GetCommentsFromSubjectRequest) (resp *types.GetCommentsFromSubjectResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
