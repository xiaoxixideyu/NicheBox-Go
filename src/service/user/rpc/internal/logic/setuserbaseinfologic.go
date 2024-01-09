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

type SetUserBaseInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetUserBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserBaseInfoLogic {
	return &SetUserBaseInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetUserBaseInfoLogic) SetUserBaseInfo(in *user.SetUserBaseInfoRequest) (*user.SetUserBaseInfoResponse, error) {
	userModel, err := l.svcCtx.UserInterface.GerUserByUid(in.Uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}

	userModel.Username = in.Username
	userModel.Introduction = in.Introduction

	err = l.svcCtx.UserInterface.UpdateUserTX(userModel)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &user.SetUserBaseInfoResponse{}, nil
}
