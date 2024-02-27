package logic

import (
	"context"
	"errors"

	"nichebox/service/box_info/model"
	"nichebox/service/box_info/rpc/internal/svc"
	"nichebox/service/box_info/rpc/pb/boxinfo"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type GetBoxInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBoxInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBoxInfoLogic {
	return &GetBoxInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBoxInfoLogic) GetBoxInfo(in *boxinfo.GetBoxInfoRequest) (*boxinfo.GetBoxInfoResponse, error) {
	box := &model.Box{
		Bid: in.Bid,
	}

	err := l.svcCtx.BoxInterface.GetBoxInfo(box)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &boxinfo.GetBoxInfoResponse{
		Name:         box.Name,
		Introduction: box.Introduction,
	}, nil
}
