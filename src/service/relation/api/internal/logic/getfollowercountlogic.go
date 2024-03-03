package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/service/relation/api/internal/svc"
	"nichebox/service/relation/api/internal/types"
	"nichebox/service/relation/rpc/pb/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFollowerCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerCountLogic {
	return &GetFollowerCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFollowerCountLogic) GetFollowerCount(req *types.GetFollowerCountRequest) (resp *types.GetFollowerCountResponse, err error) {
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

	in := relation.GetFollowerCountRequest{Uid: objectID}

	out, err := l.svcCtx.RelationRpc.GetFollowerCount(l.ctx, &in)
	if err != nil {
		l.Logger.Errorf("[RPC] Get follower count failed, err:", err)
		return nil, errors.New(http.StatusInternalServerError, "未知错误")
	}

	return &types.GetFollowerCountResponse{Count: int(out.Count)}, nil
}
