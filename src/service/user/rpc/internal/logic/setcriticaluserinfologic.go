package logic

import (
	"context"

	"nichebox/service/user/rpc/internal/svc"
	"nichebox/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetCriticalUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetCriticalUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetCriticalUserInfoLogic {
	return &SetCriticalUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetCriticalUserInfoLogic) SetCriticalUserInfo(in *user.SetCriticalUserInfoRequest) (*user.SetCriticalUserInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &user.SetCriticalUserInfoResponse{}, nil
}
