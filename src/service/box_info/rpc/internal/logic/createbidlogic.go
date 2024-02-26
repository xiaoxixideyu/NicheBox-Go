package logic

import (
	"context"

	"nichebox/service/box_info/rpc/internal/svc"
	"nichebox/service/box_info/rpc/pb/boxinfo"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateBidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBidLogic {
	return &CreateBidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateBidLogic) CreateBid(in *boxinfo.CreateBidRequest) (*boxinfo.CreateBidResponse, error) {
	bid, err := l.svcCtx.Leaves.Next("bid", 10, 0.75)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &boxinfo.CreateBidResponse{
		Bid: bid,
	}, nil
}
