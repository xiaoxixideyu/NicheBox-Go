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

type GetFollowersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFollowersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowersLogic {
	return &GetFollowersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFollowersLogic) GetFollowers(req *types.GetFollowersRequest) (resp *types.GetFollowersResponse, err error) {
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

	in := relation.GetFollowersRequest{
		Uid:   objectID,
		Page:  int32(req.Page),
		Size:  int32(req.Size),
		Order: req.Order,
	}

	out, err := l.svcCtx.RelationRpc.GetFollowers(l.ctx, &in)
	if err != nil {
		l.Logger.Errorf("[RPC] Get followers failed, err:", err)
		return nil, errors.New(http.StatusInternalServerError, "未知错误")
	}

	relations := make([]*types.Follower, 0, len(out.Followers))
	for _, f := range out.Followers {
		r := types.Follower{
			Fid:          f.Fid,
			Relationship: f.Relationship,
			UpdateTime:   f.UpdateTime,
		}
		relations = append(relations, &r)
	}

	return &types.GetFollowersResponse{Followers: relations}, nil
}
