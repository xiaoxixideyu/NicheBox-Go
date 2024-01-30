package model

import "context"

type CommentInterface interface {
	FirstOrCreateSubject(subject *Subject) error
	AddCommentAndUpdateSubjectTX(subject *Subject, comment *Comment, content *CommentContent) error
	GetComment(commentID int64) (*Comment, error)
	GetCommentContent(commentID int64) (*CommentContent, error)
	DeleteCommentAndUpdateSubjectTX(commentID int64) (*Comment, error)
}

type CommentCacheInterface interface {
	GetCommentCtx(ctx context.Context, commentID int64) (string, error)
	DeleteSubjectInfoCtx(ctx context.Context, subjectID int64) error
	DeleteCommentsBySubjectIDCtx(ctx context.Context, subjectID int64) error
	DeleteInnerFloorCommentsByRootIDCtx(ctx context.Context, rootID int64) error
}
