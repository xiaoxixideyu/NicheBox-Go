package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type CommentInterface interface {
	FirstOrCreateSubject(subject *Subject) error
	UpdateCommentLikeCount(commentID int64, delta int) (*Comment, error)
	GetSubjectBySubjectID(subjectID int64) (*Subject, error)
	AddCommentAndUpdateSubjectTX(subject *Subject, comment *Comment, content *CommentContent) error
	GetComment(commentID int64) (*Comment, error)
	GetCommentContent(commentID int64) (*CommentContent, error)
	GetRootCommentsBySubjectID(subjectID int64, page, size int, order string) ([]*Comment, error)
	BatchGetComments(ids []int64) ([]*Comment, error)
	BatchGetInnerFloorComments(rootIDs []int64, page, size int) ([]*Comment, error)
	BatchGetAllInnerFloorCommentsAndInnerFloorCounts(rootIDs []int64) ([]*Comment, []int, error)
	BatchGetAllInnerFloorCommentIDsCreateTimesAndInnerFloorCounts(rootIDs []int64) ([]*Comment, []int, error)
	BatchGetCommentsContents(ids []int64) ([]*CommentContent, error)
	GetInnerFloorCommentsAndContentsByRootID(rootID int64, page, size int) ([]*Comment, []*CommentContent, error)
	DeleteCommentAndUpdateSubjectTX(commentID int64) (*Comment, error)
}

type CommentCacheInterface interface {
	GetSubjectInfoByMessageCtx(ctx context.Context, messageID int64, messageType int) (string, error)
	GetSubjectInfoBySubjectIDCtx(ctx context.Context, subjectID int64) (string, error)
	GetCommentIndexesWithScoreBySubjectIDCtx(ctx context.Context, subjectID int64, page, size int, order string) ([]redis.Pair, error)
	GetCommentCtx(ctx context.Context, commentID int64) (string, error)
	BatchGetCommentsByIDsCtx(ctx context.Context, ids []string) (map[string]string, []string, error)
	SetCommentIndexesWithScoreBySubjectIDCtx(ctx context.Context, subjectID int64, comments []*Comment) error
	GetInnerFloorCommentIDs(ctx context.Context, rootID string, start, stop int) ([]string, error)
	SetInnerFloorCommentIDs(ctx context.Context, rootID string, comments []*Comment) error
	SetSubjectInfoCtx(ctx context.Context, subject *Subject) error
	DeleteSubjectInfoCtx(ctx context.Context, subjectID int64) error
	DeleteCommentCtx(ctx context.Context, commentID int64) error
	DeleteCommentsBySubjectIDCtx(ctx context.Context, subjectID int64) error
	DeleteInnerFloorCommentsByRootIDCtx(ctx context.Context, rootID int64) error
	BatchSetCommentsCtx(ctx context.Context, caches []*CommentCache) error
}
