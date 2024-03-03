package mqs

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"nichebox/service/relation/model"
	"nichebox/service/relation/model/dto"
	"nichebox/service/relation/rpc/internal/svc"
)

type RebuildRelationCache struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRebuildRelationCache(ctx context.Context, svcCtx *svc.ServiceContext) *RebuildRelationCache {
	return &RebuildRelationCache{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RebuildRelationCache) Consume(key, value string) error {
	msg := dto.RebuildRelationCacheMessage{}
	err := json.Unmarshal([]byte(value), &msg)
	if err != nil {
		l.Logger.Errorf("[Json][Consumer] Json unmarshal failed, err:", err)
		return err
	}
	// relationships
	followings, err := l.svcCtx.RelationInterface.GetFollowings(msg.Uid, 1, -1, "")
	if err != nil {
		l.Logger.Errorf("[MySQL][Consumer] Get followings failed, err:", err)
		return err
	}
	attrs := make([]*model.CacheRelationshipAttribute, 0, len(followings))
	for _, f := range followings {
		attr := model.CacheRelationshipAttribute{
			Fid:          f.Fid,
			Relationship: model.ConvertRelationNumberToString(f.Relationship),
			UpdateTime:   f.UpdatedAt,
		}
		attrs = append(attrs, &attr)
	}
	if len(attrs) > 0 {
		err = l.svcCtx.RelationCacheInterface.BatchAddRelationshipsCtx(l.ctx, msg.Uid, attrs)
		if err != nil {
			l.Logger.Errorf("[Redis][Consumer] Batch add relationships failed, err:", err)
			return err
		}
	}

	// count
	followingCount, err := l.svcCtx.RelationInterface.GetFollowingCount(msg.Uid)
	if err != nil {
		l.Logger.Errorf("[MySQL][Consumer] Get following count failed, err:", err)
		return err
	}
	followerCount, err := l.svcCtx.RelationInterface.GetFollowerCount(msg.Uid)
	if err != nil {
		l.Logger.Errorf("[MySQL][Consumer] Get follower count failed, err:", err)
		return err
	}
	err = l.svcCtx.RelationCacheInterface.SetRelationCountCtx(l.ctx, msg.Uid, followerCount, followingCount)
	if err != nil {
		l.Logger.Errorf("[Redis][Consumer] Batch add relationships failed, err:", err)
		return err
	}
	return nil
}
