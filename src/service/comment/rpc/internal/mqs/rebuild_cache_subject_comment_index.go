package mqs

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"nichebox/common/biz"
	"nichebox/service/comment/model/dto"
	"nichebox/service/comment/rpc/internal/svc"
	"time"
)

type RebuildCacheSubjectCommentIndex struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	lastRebuildDateMap map[int64]time.Time
}

func NewRebuildCacheSubjectCommentIndex(ctx context.Context, svcCtx *svc.ServiceContext) *RebuildCacheSubjectCommentIndex {
	return &RebuildCacheSubjectCommentIndex{
		ctx:                ctx,
		svcCtx:             svcCtx,
		Logger:             logx.WithContext(ctx),
		lastRebuildDateMap: map[int64]time.Time{},
	}
}

func (l *RebuildCacheSubjectCommentIndex) Consume(key, value string) error {
	msg := dto.RebuildCacheSubjectCommentIndexMessage{}
	err := json.Unmarshal([]byte(value), &msg)
	if err != nil {
		l.Logger.Errorf("[Json][Consumer] Json unmarshal error", err)
		return err
	}
	lastRebuildDate, ok := l.lastRebuildDateMap[msg.SubjectID]
	isNewMsg := false
	if !ok {
		isNewMsg = true
	} else {
		isNewMsg = msg.CreateDate.After(lastRebuildDate)
	}

	if !isNewMsg {
		return nil
	}

	rootComments, err := l.svcCtx.CommentInterface.GetRootCommentsBySubjectID(msg.SubjectID, 1, -1, biz.OrderByCreateTimeAsc)

	if err != nil {
		l.Logger.Errorf("[MySQL][Consumer] Get root comments by subject id error", err)
		return err
	}

	err = l.svcCtx.CommentCacheInterface.SetCommentIndexesWithScoreBySubjectIDCtx(l.ctx, msg.SubjectID, rootComments)
	if err != nil {
		l.Logger.Errorf("[Redis][Consumer] Set comment indexes zset by subject id error", err)
		return err
	}
	l.lastRebuildDateMap[msg.SubjectID] = time.Now()
	return nil
}
