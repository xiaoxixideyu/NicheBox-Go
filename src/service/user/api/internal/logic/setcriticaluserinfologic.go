package logic

import (
	"context"

	"nichebox/service/user/api/internal/svc"
	"nichebox/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetCriticalUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetCriticalUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetCriticalUserInfoLogic {
	return &SetCriticalUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetCriticalUserInfoLogic) SetCriticalUserInfo(req *types.SetCriticalUserInfoRequest) (resp *types.SetCriticalUserInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
