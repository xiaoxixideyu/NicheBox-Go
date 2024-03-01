package logic

import (
	"context"
	"errors"

	"nichebox/service/box_user/model"
	"nichebox/service/box_user/rpc/internal/svc"
	"nichebox/service/box_user/rpc/pb/boxuser"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleLogic {
	return &GetRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleLogic) GetRole(in *boxuser.GetRoleRequest) (*boxuser.GetRoleResponse, error) {
	boxUserCache, err := l.svcCtx.BoxUserCaCheInterface.GetBoxUser(in.Bid, in.Uid, l.svcCtx.Config.CacheExpire.BoxUserExist)
	if err == nil {
		return &boxuser.GetRoleResponse{
			Exist: boxUserCache.Exist,
			Role:  boxuser.UserRole(boxUserCache.Role),
		}, nil
	}

	if !errors.Is(err, redis.ErrEmptyKey) {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	boxUser, err := l.svcCtx.BoxUserInterface.GetBoxUser(in.Bid, in.Uid)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	exist := boxUser != nil

	l.svcCtx.BoxUserCaCheInterface.SetBoxUser(&model.BoxUserCache{
		Bid:   in.Bid,
		Uid:   in.Uid,
		Exist: exist,
		Role:  boxUser.Role,
	}, l.svcCtx.Config.CacheExpire.BoxUserExist)

	return &boxuser.GetRoleResponse{
		Exist: true,
		Role:  boxuser.UserRole(boxUser.Role),
	}, nil
}
