package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"nichebox/common/biz"
	"nichebox/service/box-content/rpc/boxcontent"

	"nichebox/service/box-content/api/internal/svc"
	"nichebox/service/box-content/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostListLogic {
	return &GetPostListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostListLogic) GetPostList(req *types.GetPostListRequest) (resp *types.GetPostListResponse, err error) {
	_, err = l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}
	// todo: check是不是游客的uid

	if !biz.CheckIfBoxContentOrderValid(req.Order) {
		return nil, errors.New(http.StatusBadRequest, "order无效")
	}

	in := boxcontent.GetPostListRequest{
		BoxID: req.BoxID,
		Page:  int32(req.Page),
		Size:  int32(req.Size),
		Order: req.Order,
	}
	out, err := l.svcCtx.BoxContentRpc.GetPostList(l.ctx, &in)
	if err != nil {
		l.Logger.Errorf("[Rpc] Get post list error", err)
		rpcStatus, ok := status.FromError(err)
		if ok {
			if rpcStatus.Code() == codes.OutOfRange {
				return nil, errors.New(http.StatusBadRequest, err.Error())
			}
		}
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	return &types.GetPostListResponse{IDs: out.IDs}, nil
}
