package logic

import (
	"context"

	"nichebox/service/relation/rpc/internal/svc"
	"nichebox/service/relation/rpc/pb/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowLogic) Follow(in *relation.FollowRequest) (*relation.FollowResponse, error) {
	err := l.svcCtx.RelationInterface.AddFollow(in.Uid, in.Fid)
	if err != nil {
		l.Logger.Errorf("[MySQL] Add follow failed, err:", err)
		return nil, err
	}
	l.svcCtx.RelationCacheInterface.DeleteRelationCountCtx(context.Background(), in.Uid)
	l.svcCtx.RelationCacheInterface.RemoveRelationshipsCtx(context.Background(), in.Uid)

	return &relation.FollowResponse{}, nil
}
