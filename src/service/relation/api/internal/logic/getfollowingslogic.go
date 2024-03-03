package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/common/biz"
	"nichebox/service/relation/rpc/pb/relation"

	"nichebox/service/relation/api/internal/svc"
	"nichebox/service/relation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowingsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFollowingsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingsLogic {
	return &GetFollowingsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFollowingsLogic) GetFollowings(req *types.GetFollowingsRequest) (resp *types.GetFollowingsResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	if biz.CheckIfRelationOrderValid(req.Order) {
		return nil, errors.New(http.StatusBadRequest, "invalid order")
	}

	// objectID represents the user who will be queried
	objectID := req.Uid
	if objectID != uid {
		// uid queries other users
		// todo: check 隐私

		// user can just query fixed page, size and order
		req.Page = 1
		req.Size = 10
		req.Order = biz.OrderByCreateTimeDesc
	}

	in := relation.GetFollowingsRequest{
		Uid:   objectID,
		Page:  int32(req.Page),
		Size:  int32(req.Size),
		Order: req.Order,
	}

	out, err := l.svcCtx.RelationRpc.GetFollowings(l.ctx, &in)
	if err != nil {
		l.Logger.Errorf("[RPC] Get followers failed, err:", err)
		return nil, errors.New(http.StatusInternalServerError, "未知错误")
	}

	relations := make([]*types.Following, 0, len(out.Followings))
	for _, f := range out.Followings {
		r := types.Following{
			Fid:          f.Fid,
			Relationship: f.Relationship,
			UpdateTime:   f.UpdateTime,
		}
		relations = append(relations, &r)
	}

	return &types.GetFollowingsResponse{Followings: relations}, nil
}
