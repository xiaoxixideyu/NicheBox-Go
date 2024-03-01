package logic

import (
	"context"
	"errors"

	"nichebox/service/box_user/model"
	"nichebox/service/box_user/rpc/internal/svc"
	"nichebox/service/box_user/rpc/pb/boxuser"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AddBoxUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddBoxUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBoxUserLogic {
	return &AddBoxUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddBoxUserLogic) AddBoxUser(in *boxuser.AddBoxUserRequest) (*boxuser.AddBoxUserResponse, error) {
	boxUser := &model.BoxUser{
		Bid:  in.Bid,
		Uid:  in.Uid,
		Role: int(boxuser.UserRole_Member),
	}

	if err := l.svcCtx.BoxUserInterface.AddBoxUser(boxUser); err != nil {
		if errors.Is(err, model.ErrBoxUserExisted) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}

	l.svcCtx.BoxUserCaCheInterface.SetBoxUser(&model.BoxUserCache{
		Bid:   boxUser.Bid,
		Uid:   boxUser.Uid,
		Exist: true,
		Role:  boxUser.Role,
	}, l.svcCtx.Config.CacheExpire.BoxUserExist)

	return &boxuser.AddBoxUserResponse{}, nil
}
