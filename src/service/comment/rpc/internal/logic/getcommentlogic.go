package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"nichebox/service/comment/model"
	"nichebox/service/comment/rpc/internal/svc"
	"nichebox/service/comment/rpc/pb/comment"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLogic {
	return &GetCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentLogic) GetComment(in *comment.GetCommentRequest) (*comment.GetCommentResponse, error) {
	commentStr, err := l.svcCtx.CommentCacheInterface.GetCommentCtx(l.ctx, in.CommentID)
	needQueryDB := false
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			l.Logger.Errorf("[Redis] Get comment error", err)
		}
		needQueryDB = true
	}

	// try to use comment cache
	if !needQueryDB {
		var commentInfoCache model.CommentCache
		err = json.Unmarshal([]byte(commentStr), &commentInfoCache)
		if err != nil {
			// cache data invalid, delete cache and query db
			l.Logger.Errorf("[Json] Unmarshal comment info cache error", err)
			// todo: delete cache
			needQueryDB = true
		} else {
			// cache valid, return comment from cache
			info := comment.CommentInfo{
				CommentID:          commentInfoCache.CommentID,
				SubjectID:          commentInfoCache.SubjectID,
				RootID:             commentInfoCache.RootID,
				ParentID:           commentInfoCache.ParentID,
				DialogID:           commentInfoCache.DialogID,
				OwnerID:            commentInfoCache.OwnerID,
				ThumbsUp:           false,
				Floor:              int32(commentInfoCache.Floor),
				CreateTime:         commentInfoCache.CreatedAt.Format("2006-01-02"),
				InnerFloorCount:    int32(commentInfoCache.InnerFloorCount),
				InnerFloorComments: nil, // todo: inner floor content
				Content:            commentInfoCache.Content,
			}
			return &comment.GetCommentResponse{Comment: &info}, nil
		}
	}

	// need to query DB
	//if needQueryDB {
	cmt, err := l.svcCtx.CommentInterface.GetComment(in.CommentID)
	if err != nil {
		l.Logger.Errorf("[MySQL] Get Comment error", err)
		return nil, err
	}
	content, err := l.svcCtx.CommentInterface.GetCommentContent(in.CommentID)
	if err != nil {
		l.Logger.Errorf("[MySQL] Get comment content error", err)
		return nil, err
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
		CreateTime:         cmt.CreatedAt.Format(time.DateTime),
		InnerFloorCount:    int32(cmt.InnerFloorCount),
		InnerFloorComments: nil,
		Content:            content.Content,
	}
	return &comment.GetCommentResponse{Comment: &info}, nil
	//}
	//return nil, status.Error(codes.Unknown, "unknown error")
}
