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

type GetFollowerCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerCountLogic {
	return &GetFollowerCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowerCountLogic) GetFollowerCount(in *relation.GetFollowerCountRequest) (*relation.GetFollowerCountResponse, error) {
	needQueryDB := false
	count := 0

	followers, _, err := l.svcCtx.RelationCacheInterface.GetRelationCountCtx(l.ctx, in.Uid)
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
		count = followers
	}

	if needQueryDB {
		count, err = l.svcCtx.RelationInterface.GetFollowerCount(in.Uid)
		if err != nil {
			l.Logger.Errorf("[MySQL] Get follower count failed, err:", err)
			return nil, err
		}
	}

	return &relation.GetFollowerCountResponse{Count: int32(count)}, nil
}
