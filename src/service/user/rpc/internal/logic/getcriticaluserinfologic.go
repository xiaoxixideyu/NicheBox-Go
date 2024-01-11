package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"nichebox/service/user/rpc/internal/svc"
	"nichebox/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCriticalUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCriticalUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCriticalUserInfoLogic {
	return &GetCriticalUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCriticalUserInfoLogic) GetCriticalUserInfo(in *user.GetCriticalUserInfoRequest) (*user.GetCriticalUserInfoResponse, error) {
	userModel, err := l.svcCtx.UserInterface.GetUserByUid(in.Uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &user.GetCriticalUserInfoResponse{Email: userModel.Email, Telephone: userModel.Telephone}, nil
}
