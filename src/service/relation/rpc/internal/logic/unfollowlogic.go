package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"nichebox/service/relation/rpc/internal/svc"
	"nichebox/service/relation/rpc/pb/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnfollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnfollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfollowLogic {
	return &UnfollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnfollowLogic) Unfollow(in *relation.UnfollowRequest) (*relation.UnfollowResponse, error) {
	err := l.svcCtx.RelationInterface.RemoveFollow(in.Uid, in.Fid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "您尚未关注这个用户")
		}
		l.Logger.Errorf("[MySQL] Remove follow failed, err:", err)
		return nil, err
	}
	l.svcCtx.RelationCacheInterface.DeleteRelationCountCtx(context.Background(), in.Uid)
	l.svcCtx.RelationCacheInterface.RemoveRelationshipsCtx(context.Background(), in.Uid)

	return &relation.UnfollowResponse{}, nil
}
