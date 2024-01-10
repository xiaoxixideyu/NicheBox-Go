package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nichebox/common/cryptx"
	redisBiz "nichebox/service/user/model/redis"

	"nichebox/service/user/rpc/internal/svc"
	"nichebox/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ForgetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewForgetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForgetPasswordLogic {
	return &ForgetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ForgetPasswordLogic) ForgetPassword(in *user.ForgetPasswordRequest) (*user.ForgetPasswordResponse, error) {
	key := redisBiz.KeyPrefixUser + redisBiz.KeyForgetPasswordCode + in.Email
	val, err := l.svcCtx.UserRedisInterface.GetVerificationCode(l.ctx, key)
	if err != nil && errors.As(err, &redis.ErrEmptyKey) {
		return nil, status.Error(codes.NotFound, "验证码错误")
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if val != in.Code {
		return nil, status.Error(codes.NotFound, "验证码错误")
	}

	err = l.svcCtx.UserInterface.UpdatePasswordByEmail(in.Email, cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.NewPassword))
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &user.ForgetPasswordResponse{}, nil
}
