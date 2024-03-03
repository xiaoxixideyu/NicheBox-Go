package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"nichebox/service/relation/model/dto"

	"nichebox/service/relation/rpc/internal/svc"
	"nichebox/service/relation/rpc/pb/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowingCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowingCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingCountLogic {
	return &GetFollowingCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowingCountLogic) GetFollowingCount(in *relation.GetFollowingCountRequest) (*relation.GetFollowingCountResponse, error) {
	needQueryDB := false
	count := 0

	_, followings, err := l.svcCtx.RelationCacheInterface.GetRelationCountCtx(l.ctx, in.Uid)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			// rebuild cache
			msg := dto.RebuildRelationCacheMessage{Uid: in.Uid}
			bytes, _ := json.Marshal(msg)
			err := l.svcCtx.KqRebuildRelationCachePusher.Push(string(bytes))
			if err != nil {
				l.Logger.Errorf("[Producer] Push rebuild relation cache failed, err:", err)
			}
		} else {
			l.Logger.Errorf("[Redis] Get relationship count failed, err:", err)
		}
		needQueryDB = true
	} else {
		count = followings
	}

	if needQueryDB {
		count, err = l.svcCtx.RelationInterface.GetFollowingCount(in.Uid)
		if err != nil {
			l.Logger.Errorf("[MySQL] Get following count failed, err:", err)
			return nil, err
		}
	}

	return &relation.GetFollowingCountResponse{Count: int32(count)}, nil
}
