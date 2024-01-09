package logic

import (
	"context"

	"nichebox/service/user/api/internal/svc"
	"nichebox/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckVerificationCodeCriticalUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckVerificationCodeCriticalUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckVerificationCodeCriticalUserInfoLogic {
	return &CheckVerificationCodeCriticalUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckVerificationCodeCriticalUserInfoLogic) CheckVerificationCodeCriticalUserInfo(req *types.CheckVerificationCodeCriticalUserInfoRequest) (resp *types.CheckVerificationCodeCriticalUserInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
