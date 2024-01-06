package logic

import (
	"context"

	"nichebox/service/user/api/internal/svc"
	"nichebox/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckEmailExistsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckEmailExistsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckEmailExistsLogic {
	return &CheckEmailExistsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckEmailExistsLogic) CheckEmailExists(req *types.CheckEmailExistsReqeust) (resp *types.CheckEmailExistsResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
