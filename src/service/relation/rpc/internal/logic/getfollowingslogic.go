package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"math"
	"nichebox/common/biz"
	"nichebox/service/relation/model"
	"nichebox/service/relation/model/dto"
	"sort"
	"time"

	"nichebox/service/relation/rpc/internal/svc"
	"nichebox/service/relation/rpc/pb/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowingsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowingsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingsLogic {
	return &GetFollowingsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowingsLogic) GetFollowings(in *relation.GetFollowingsRequest) (*relation.GetFollowingsResponse, error) {
	var relations []*relation.RelationMessage
	needQueryDB := false

	caches, err := l.svcCtx.RelationCacheInterface.GetAllRelationshipsCtx(l.ctx, in.Uid)
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
			l.Logger.Errorf("[Redis] Get all relationships failed, err:", err)
		}
		needQueryDB = true
	}
	// use cache
	if !needQueryDB {
		relations = make([]*relation.RelationMessage, 0, len(caches))
		for _, cache := range caches {
			if cache.Relationship == model.ConvertRelationNumberToString(model.RelationFollow) || cache.Relationship == model.ConvertRelationNumberToString(model.RelationFriend) {
				r := relation.RelationMessage{
					Fid:          cache.Fid,
					Relationship: cache.Relationship,
					UpdateTime:   cache.UpdateTime.Format(time.DateTime),
				}
				relations = append(relations, &r)
			}
		}
		if in.Size != -1 {
			offset := int((in.Page - 1) * in.Size)
			size := math.Min(float64(offset+int(in.Size)), float64(len(relations)))
			if in.Order == biz.OrderByCreateTimeAsc {
				sort.Slice(relations, func(i, j int) bool {
					return relations[i].UpdateTime > relations[j].UpdateTime
				})
			} else {
				sort.Slice(relations, func(i, j int) bool {
					return relations[i].UpdateTime < relations[j].UpdateTime
				})
			}
			relations = relations[offset:int(size)]
		}
	}

	if needQueryDB {
		fs, err := l.svcCtx.RelationInterface.GetFollowings(in.Uid, int(in.Page), int(in.Size), in.Order)
		if err != nil {
			l.Logger.Errorf("[MySQL] Get followings failed, err:", err)
			return nil, err
		}
		relations = make([]*relation.RelationMessage, 0, len(fs))
		for _, f := range fs {
			r := relation.RelationMessage{
				Fid:          f.Fid,
				Relationship: model.ConvertRelationNumberToString(f.Relationship),
				UpdateTime:   f.UpdatedAt.Format(time.DateTime),
			}
			relations = append(relations, &r)
		}
	}

	return &relation.GetFollowingsResponse{Followings: relations}, nil
}
