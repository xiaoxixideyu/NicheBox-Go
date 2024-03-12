package mqs

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"nichebox/service/comment/model"
	"nichebox/service/comment/model/dto"
	"nichebox/service/comment/rpc/internal/svc"
	"strconv"
)

type RebuildCacheInnerFloorCommentIndex struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRebuildCacheInnerFloorCommentIndex(ctx context.Context, svcCtx *svc.ServiceContext) *RebuildCacheInnerFloorCommentIndex {
	return &RebuildCacheInnerFloorCommentIndex{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RebuildCacheInnerFloorCommentIndex) Consume(key, value string) error {
	msg := dto.RebuildCacheInnerFloorCommentIndexMessage{}
	err := json.Unmarshal([]byte(value), &msg)
	if err != nil {
		l.Logger.Errorf("[Json][Consumer] Json unmarshal error", err)
		return err
	}

	//subComments, innerFloorCounts, err := l.svcCtx.CommentInterface.BatchGetAllInnerFloorCommentsAndInnerFloorCounts(msg.RootIDs)
	subComments, innerFloorCounts, err := l.svcCtx.CommentInterface.BatchGetAllInnerFloorCommentIDsCreateTimesAndInnerFloorCounts(msg.RootIDs)

	offset := 0
	for i := 0; i < len(msg.RootIDs); i++ {
		count := innerFloorCounts[i]
		if count == 0 {
			continue
		}
		innerComments := make([]*model.Comment, 0, count)
		j := offset
		for ; j < offset+count; j++ {
			innerComments = append(innerComments, subComments[j])
		}
		err := l.svcCtx.CommentCacheInterface.SetInnerFloorCommentIDs(l.ctx, strconv.FormatInt(msg.RootIDs[i], 10), innerComments)
		if err != nil {
			l.Logger.Errorf("[Redis][Consumer] Set inner floor comment index error", err)
		}
		offset = j
	}
	return nil
}
