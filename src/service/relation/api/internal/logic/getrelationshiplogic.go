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

type GetRelationshipLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRelationshipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRelationshipLogic {
	return &GetRelationshipLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRelationshipLogic) GetRelationship(req *types.GetRelationshipRequest) (resp *types.GetRelationshipResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	in := relation.GetRelationshipRequest{
		Uid: uid,
		Fid: req.Fid,
	}
	out, err := l.svcCtx.RelationRpc.GetRelationship(l.ctx, &in)
	if err != nil {
		l.Logger.Errorf("[RPC] Get relationship failed, err:", err)
		return nil, errors.New(http.StatusInternalServerError, "未知错误")
	}

	return &types.GetRelationshipResponse{Relationship: out.Relationship}, nil
}
