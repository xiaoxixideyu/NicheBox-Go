package logic

import (
	"context"

	"nichebox/service/user/api/internal/svc"
	"nichebox/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendVerificationCodeCriticalUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendVerificationCodeCriticalUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendVerificationCodeCriticalUserInfoLogic {
	return &SendVerificationCodeCriticalUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendVerificationCodeCriticalUserInfoLogic) SendVerificationCodeCriticalUserInfo(req *types.SendVerificationCodeCriticalUserInfoRequest) (resp *types.SendVerificationCodeCriticalUserInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
