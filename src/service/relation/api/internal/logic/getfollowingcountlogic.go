package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/service/relation/rpc/pb/relation"

	"nichebox/service/relation/api/internal/svc"
	"nichebox/service/relation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowingCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFollowingCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingCountLogic {
	return &GetFollowingCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFollowingCountLogic) GetFollowingCount(req *types.GetFollowingCountRequest) (resp *types.GetFollowingCountResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	// objectID represents the user who will be queried
	objectID := req.Uid
	if objectID != uid {
		// uid queries other users
		// todo: check 隐私
	}

	in := relation.GetFollowingCountRequest{Uid: objectID}

	out, err := l.svcCtx.RelationRpc.GetFollowingCount(l.ctx, &in)
	if err != nil {
		l.Logger.Errorf("[RPC] Get following count failed, err:", err)
		return nil, errors.New(http.StatusInternalServerError, "未知错误")
	}

	return &types.GetFollowingCountResponse{Count: int(out.Count)}, nil
}
