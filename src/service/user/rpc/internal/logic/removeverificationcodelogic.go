package logic

import (
	"context"

	"nichebox/service/user/rpc/internal/svc"
	"nichebox/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveVerificationCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveVerificationCodeLogic {
	return &RemoveVerificationCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveVerificationCodeLogic) RemoveVerificationCode(in *user.RemoveVerificationCodeRequest) (*user.RemoveVerificationCodeResponse, error) {
	err := l.svcCtx.UserCacheInterface.RemoveVerificationCode(l.ctx, in.Key)
	if err != nil {
		return nil, err
	}

	return &user.RemoveVerificationCodeResponse{}, nil
}
