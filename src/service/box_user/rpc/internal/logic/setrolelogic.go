package logic

import (
	"context"

	"nichebox/service/box_user/model"
	"nichebox/service/box_user/rpc/internal/common"
	"nichebox/service/box_user/rpc/internal/svc"
	"nichebox/service/box_user/rpc/pb/boxuser"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SetRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRoleLogic {
	return &SetRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetRoleLogic) SetRole(in *boxuser.SetRoleRequest) (*boxuser.SetRoleResponse, error) {
	boxUser := &model.BoxUser{
		Bid:  in.Bid,
		Uid:  in.Uid,
		Role: int(in.Role),
	}

	if err := l.svcCtx.BoxUserInterface.UpdateRole(boxUser); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	common.PushRemoveCacheBoxUserExist(boxUser.Bid, boxUser.Uid, l.svcCtx)

	return &boxuser.SetRoleResponse{}, nil
}
