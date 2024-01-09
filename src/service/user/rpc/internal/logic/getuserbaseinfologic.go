package logic

import (
	"context"
	"errors"

	"nichebox/service/user/rpc/internal/svc"
	"nichebox/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type GetUserBaseInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserBaseInfoLogic {
	return &GetUserBaseInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserBaseInfoLogic) GetUserBaseInfo(in *user.GetUserBaseInfoRequest) (*user.GetUserBaseInfoResponse, error) {
	userModel, err := l.svcCtx.UserInterface.GetUserByUid(in.Uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &user.GetUserBaseInfoResponse{
		Username:     userModel.Username,
		Introduction: userModel.Introduction,
	}, nil
}
