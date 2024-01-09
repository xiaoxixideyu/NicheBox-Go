package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"nichebox/common/cryptx"

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
	userModel, err := l.svcCtx.UserInterface.GetUserByUid(in.Uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}

	userModel.Password = cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	err = l.svcCtx.UserInterface.UpdateUserTX(userModel)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &user.SetCriticalUserInfoResponse{}, nil
}
