package logic

import (
	"context"
	"net/http"
	"strconv"

	"nichebox/service/box_info/api/internal/svc"
	"nichebox/service/box_info/api/internal/types"
	"nichebox/service/box_info/rpc/pb/boxinfo"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetBoxInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBoxInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBoxInfoLogic {
	return &GetBoxInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBoxInfoLogic) GetBoxInfo(req *types.GetBoxInfoRequest) (resp *types.GetBoxInfoResponse, err error) {
	bid, err := strconv.ParseInt(req.BoxId, 10, 64)
	if err != nil {
		return nil, errors.New(http.StatusBadRequest, "box_id 无效")
	}

	res, err := l.svcCtx.BoxInfoRpc.GetBoxInfo(l.ctx, &boxinfo.GetBoxInfoRequest{
		Bid: bid,
	})
	if err != nil {
		rpcStatus, ok := status.FromError(err)
		if ok {
			if rpcStatus.Code() == codes.NotFound {
				return nil, errors.New(http.StatusNotFound, "盒子未找到")
			}
		}
		return nil, errors.New(http.StatusInternalServerError, "get box info 未知错误: 1")
	}

	return &types.GetBoxInfoResponse{
		Name:         res.Name,
		Introduction: res.Introduction,
	}, nil
}
