package mqs

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"nichebox/service/comment/model/dto"
	"nichebox/service/comment/rpc/internal/svc"
)

type UpdateCommentLikeCount struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentLikeCount(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentLikeCount {
	return &UpdateCommentLikeCount{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCommentLikeCount) Consume(key, value string) error {
	msg := dto.UpdateCommentLikeCountMessage{}
	err := json.Unmarshal([]byte(value), &msg)
	if err != nil {
		l.Logger.Errorf("[Json][Consumer] Json unmarshal error", err)
		return err
	}
	comment, err := l.svcCtx.CommentInterface.UpdateCommentLikeCount(msg.CommentID, msg.Delta)
	if err != nil {
		l.Logger.Errorf("[MySQL][Consumer] Update comment like count error", err)
		// todo: retry or save db
		return err
	}
	l.svcCtx.CommentCacheInterface.DeleteCommentsBySubjectIDCtx(l.ctx, comment.SubjectID)
	l.svcCtx.CommentCacheInterface.DeleteCommentCtx(l.ctx, msg.CommentID)

	return nil
}
