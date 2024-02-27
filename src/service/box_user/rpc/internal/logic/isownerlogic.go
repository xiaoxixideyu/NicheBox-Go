package logic

import (
	"context"

	"nichebox/service/box_user/model"
	"nichebox/service/box_user/rpc/internal/svc"
	"nichebox/service/box_user/rpc/pb/boxuser"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IsOwnerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsOwnerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsOwnerLogic {
	return &IsOwnerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsOwnerLogic) IsOwner(in *boxuser.IsOwnerRequest) (*boxuser.IsOwnerResponse, error) {
	boxUser := &model.BoxUser{
		Bid:  in.Bid,
		Uid:  in.Uid,
		Role: int(boxuser.UserRole_Owner),
	}

	exists, err := l.svcCtx.BoxUserInterface.IsOwnerExists(boxUser)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &boxuser.IsOwnerResponse{
		Exists: exists,
	}, nil
}
