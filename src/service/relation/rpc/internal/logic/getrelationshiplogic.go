package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"nichebox/service/relation/model"
	"nichebox/service/relation/model/dto"

	"nichebox/service/relation/rpc/internal/svc"
	"nichebox/service/relation/rpc/pb/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRelationshipLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRelationshipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRelationshipLogic {
	return &GetRelationshipLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRelationshipLogic) GetRelationship(in *relation.GetRelationshipRequest) (*relation.GetRelationshipResponse, error) {
	attr, err := l.svcCtx.RelationCacheInterface.GetRelationshipCtx(l.ctx, in.Uid, in.Fid)
	needQueryDB := false
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
			l.Logger.Errorf("[Redis] Get relationship failed, err:", err)
		}
		needQueryDB = true
	}
	// use cache
	if !needQueryDB {
		if attr.Relationship == model.ConvertRelationNumberToString(model.RelationNone) || attr.Relationship == model.ConvertRelationNumberToString(model.RelationFriend) {
			return &relation.GetRelationshipResponse{Relationship: attr.Relationship}, nil
		}
		if attr.Fid == in.Fid {
			return &relation.GetRelationshipResponse{Relationship: attr.Relationship}, nil
		}
		// if cache result from fid's hash, we might need to convert the relationship
		if attr.Fid == in.Uid {
			var relationship string
			if attr.Relationship == model.ConvertRelationNumberToString(model.RelationFollow) {
				relationship = model.ConvertRelationNumberToString(model.RelationFollower)
			} else {
				relationship = model.ConvertRelationNumberToString(model.RelationFollow)
			}
			return &relation.GetRelationshipResponse{Relationship: relationship}, nil
		}
	}

	// query DB
	needQueryFid := false
	var relationshipStr string
	relationship, err := l.svcCtx.RelationInterface.GetRelationship(in.Uid, in.Fid)
	if err != nil {
		// occur other errors, return error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			l.Logger.Errorf("[MySQL] Get relationship failed, err:", err)
			return nil, err
		}
		// not found
		needQueryFid = true
	}
	// use result from uid
	if !needQueryFid {
		relationshipStr = model.ConvertRelationNumberToString(relationship.Relationship)
	}
	// record not found from uid, try to query fid
	if needQueryFid {
		relationship, err := l.svcCtx.RelationInterface.GetRelationship(in.Fid, in.Uid)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				l.Logger.Errorf("[MySQL] Get relationship failed, err:", err)
				return nil, err
			}
			// not found relationship, represent there is no relationship between uid and fid, going to return relation none
			relationshipStr = model.ConvertRelationNumberToString(model.RelationNone)
		}
		// use result from fid
		if err == nil {
			if relationship.Relationship == model.RelationNone || relationship.Relationship == model.RelationFriend {
				relationshipStr = model.ConvertRelationNumberToString(relationship.Relationship)
			}
			// convert the relationship
			if relationship.Relationship == model.RelationFollow {
				relationshipStr = model.ConvertRelationNumberToString(model.RelationFollower)
			}
		}
	}

	return &relation.GetRelationshipResponse{Relationship: relationshipStr}, nil
}
