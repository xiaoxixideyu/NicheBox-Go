package logic

import (
	"context"
	"nichebox/service/relation/model"
	"time"

	"nichebox/service/relation/rpc/internal/svc"
	"nichebox/service/relation/rpc/pb/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowersLogic {
	return &GetFollowersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowersLogic) GetFollowers(in *relation.GetFollowersRequest) (*relation.GetFollowersResponse, error) {
	var relations []*relation.RelationMessage
	fs, err := l.svcCtx.RelationInterface.GetFollowers(in.Uid, int(in.Page), int(in.Size), in.Order)
	if err != nil {
		l.Logger.Errorf("[MySQL] Get followers failed, err:", err)
		return nil, err
	}
	relations = make([]*relation.RelationMessage, 0, len(fs))
	for _, f := range fs {
		r := relation.RelationMessage{
			Fid:          f.Uid,
			Relationship: model.ConvertRelationNumberToString(f.Relationship),
			UpdateTime:   f.UpdatedAt.Format(time.DateTime),
		}
		relations = append(relations, &r)
	}

	return &relation.GetFollowersResponse{Followers: relations}, nil
}
