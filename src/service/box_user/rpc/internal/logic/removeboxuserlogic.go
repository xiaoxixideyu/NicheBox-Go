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

type RemoveBoxUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveBoxUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveBoxUserLogic {
	return &RemoveBoxUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveBoxUserLogic) RemoveBoxUser(in *boxuser.RemoveBoxUserRequest) (*boxuser.RemoveBoxUserResponse, error) {
	if err := common.PushRemoveCacheBoxUserExist(in.Bid, in.Uid, l.svcCtx); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if err := l.svcCtx.BoxUserInterface.RemoveBoxUser(&model.BoxUser{
		Bid: in.Bid,
		Uid: in.Uid,
	}); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &boxuser.RemoveBoxUserResponse{}, nil
}
