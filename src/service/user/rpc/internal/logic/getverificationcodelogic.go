package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	code, err := l.svcCtx.UserCacheInterface.GetVerificationCode(l.ctx, in.Key)
	if err != nil {
		return nil, err
	}
	if err != nil && errors.As(err, &redis.ErrEmptyKey) {
		return nil, status.Error(codes.NotFound, "验证码不存在")
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &user.GetVerificationCodeResponse{Val: code}, nil
}
