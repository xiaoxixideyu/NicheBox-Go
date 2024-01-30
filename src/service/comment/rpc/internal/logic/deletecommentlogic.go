package logic

import (
	"context"
	"nichebox/service/comment/rpc/internal/svc"
	"nichebox/service/comment/rpc/pb/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentLogic) DeleteComment(in *comment.DeleteCommentRequest) (*comment.DeleteCommentResponse, error) {
	cmt, err := l.svcCtx.CommentInterface.DeleteCommentAndUpdateSubjectTX(in.CommentID)
	if err != nil {
		l.Logger.Errorf("[MySQL] Delete comment and update subject error", err)
		return nil, err
	}

	// remove cache (cache aside)
	l.svcCtx.CommentCacheInterface.DeleteSubjectInfoCtx(context.Background(), cmt.SubjectID)
	l.svcCtx.CommentCacheInterface.DeleteCommentsBySubjectIDCtx(context.Background(), cmt.SubjectID)
	if cmt.RootID != 0 {
		l.svcCtx.CommentCacheInterface.DeleteInnerFloorCommentsByRootIDCtx(context.Background(), cmt.RootID)
	}

	return &comment.DeleteCommentResponse{}, nil
}
