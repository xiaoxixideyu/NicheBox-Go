package logic

import (
	"context"
	"errors"

	"nichebox/common/cryptx"
	"nichebox/service/user/rpc/internal/svc"
	"nichebox/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type GetUidByEmailAndPwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUidByEmailAndPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUidByEmailAndPwdLogic {
	return &GetUidByEmailAndPwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUidByEmailAndPwdLogic) GetUidByEmailAndPwd(in *user.GetUidByEmailAndPwdRequest) (*user.GetUidByEmailAndPwdResponse, error) {
	res, err := l.svcCtx.UserInterface.GetUserByEmail(in.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "用户不存在")
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}

	pwd := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if pwd != res.Password {
		return nil, status.Error(codes.Unauthenticated, "密码错误")
	}
	return &user.GetUidByEmailAndPwdResponse{
		Uid: res.Uid,
	}, nil
}
