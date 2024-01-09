package logic

import (
	"context"

	"nichebox/service/user/rpc/internal/svc"
	"nichebox/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVerificationCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVerificationCodeLogic {
	return &GetVerificationCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVerificationCodeLogic) GetVerificationCode(in *user.GetVerificationCodeRequest) (*user.GetVerificationCodeResponse, error) {
	code, err := l.svcCtx.UserRedisInterface.GetVerificationCode(l.ctx, in.Key)
	if err != nil {
		return nil, err
	}

	return &user.GetVerificationCodeResponse{Val: code}, nil
}
