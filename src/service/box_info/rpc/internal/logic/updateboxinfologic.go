package logic

import (
	"context"

	"nichebox/service/box_info/model"
	"nichebox/service/box_info/rpc/internal/svc"
	"nichebox/service/box_info/rpc/pb/boxinfo"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateBoxInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBoxInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBoxInfoLogic {
	return &UpdateBoxInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBoxInfoLogic) UpdateBoxInfo(in *boxinfo.UpdateBoxInfoRequest) (*boxinfo.UpdateBoxInfoResponse, error) {
	box := &model.Box{
		Bid:          in.Bid,
		Name:         in.Name,
		Introduction: in.Introduction,
	}

	tx := l.svcCtx.BoxInterface.GetTx()
	exists, err := l.svcCtx.BoxInterface.IsBoxExistsByTx(box, tx)
	if err != nil {
		tx.Rollback()
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !exists {
		tx.Rollback()
		return nil, status.Error(codes.Aborted, "盒子不存在")
	}

	if err := l.svcCtx.BoxInterface.UpdateBoxByTx(box, tx); err != nil {
		tx.Rollback()
		return nil, status.Error(codes.Internal, err.Error())
	}

	tx.Commit()

	return &boxinfo.UpdateBoxInfoResponse{}, nil
}
