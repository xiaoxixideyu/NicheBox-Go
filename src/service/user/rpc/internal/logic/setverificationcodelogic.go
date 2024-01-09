package logic

import (
	"context"

	"nichebox/service/user/rpc/internal/svc"
	"nichebox/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetVerificationCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetVerificationCodeLogic {
	return &SetVerificationCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetVerificationCodeLogic) SetVerificationCode(in *user.SetVerificationCodeRequest) (*user.SetVerificationCodeResponse, error) {
	err := l.svcCtx.UserRedisInterface.SetVerificationCode(l.ctx, in.Key, in.Val, int(in.Expiration))
	if err != nil {
		return nil, err
	}

	return &user.SetVerificationCodeResponse{}, nil
}
