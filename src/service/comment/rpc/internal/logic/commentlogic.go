package logic

import (
	"context"
	"nichebox/service/comment/model"

	"nichebox/service/comment/rpc/internal/svc"
	"nichebox/service/comment/rpc/pb/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentLogic {
	return &CommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentLogic) Comment(in *comment.CommentRequest) (*comment.CommentResponse, error) {
	sbj := model.Subject{
		TypeID:       uint8(in.MessageType),
		MessageID:    in.MessageID,
		CommentCount: 0,
	}
	err := l.svcCtx.CommentInterface.FirstOrCreateSubject(&sbj)
	if err != nil {
		return nil, err
	}
	cmt := model.Comment{
		SubjectID:       int64(sbj.ID),
		RootID:          in.RootID,
		ParentID:        in.ParentID,
		DialogID:        in.DialogID,
		OwnerID:         in.Uid,
		InnerFloorCount: 0,
		Status:          CommentStatusNormal,
	}
	content := model.CommentContent{
		Content: in.Content,
	}
	err = l.svcCtx.CommentInterface.AddCommentAndUpdateSubjectTX(&sbj, &cmt, &content)
	if err != nil {
		return nil, err
	}
	// remove cache (cache aside)
	l.svcCtx.CommentCacheInterface.DeleteSubjectInfoCtx(context.Background(), int64(sbj.ID))
	l.svcCtx.CommentCacheInterface.DeleteCommentsBySubjectIDCtx(context.Background(), int64(sbj.ID))
	if cmt.RootID != 0 {
		l.svcCtx.CommentCacheInterface.DeleteInnerFloorCommentsByRootIDCtx(context.Background(), cmt.RootID)
	}

	info := comment.CommentInfo{
		CommentID:          cmt.CommentID,
		SubjectID:          cmt.SubjectID,
		RootID:             cmt.RootID,
		ParentID:           cmt.ParentID,
		DialogID:           cmt.DialogID,
		OwnerID:            cmt.OwnerID,
		ThumbsUp:           false,
		Floor:              int32(cmt.Floor),
		CreateTime:         cmt.CreatedAt.Format("2006-01-02"),
		InnerFloorCount:    int32(cmt.InnerFloorCount),
		InnerFloorComments: nil,
		Content:            content.Content,
	}
	return &comment.CommentResponse{Comment: &info}, nil
}
