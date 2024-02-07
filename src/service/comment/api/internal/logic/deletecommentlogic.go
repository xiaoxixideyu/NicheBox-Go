package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/common/biz"
	"nichebox/service/comment/api/internal/svc"
	"nichebox/service/comment/api/internal/types"
	"nichebox/service/comment/rpc/pb/comment"
	"nichebox/service/post/rpc/pb/post"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	OperatorContentAuthor = "author"
	OperatorAdmin         = "admin"
	OperatorCommentOwner  = "owner"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentLogic) DeleteComment(req *types.DeleteCommentRequest) (resp *types.DeleteCommentResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}
	inGetComment := comment.GetCommentRequest{CommentID: req.CommentID}
	cmt, err := l.svcCtx.CommentRpc.GetComment(l.ctx, &inGetComment)
	if err != nil {
		l.Logger.Errorf("[RPC] Get comment error", err)
		return nil, errors.New(http.StatusInternalServerError, "发生未知错误")
	}

	// check operator role
	if req.Operator == OperatorAdmin {
		// todo: check admin

	} else if req.Operator == OperatorCommentOwner {
		if uid != cmt.Comment.OwnerID {
			return nil, errors.New(http.StatusUnauthorized, "你没有权限删除本评论")
		}

	} else if req.Operator == OperatorContentAuthor {
		// get subject info
		inGetSubject := comment.GetSubjectRequest{SubjectID: cmt.Comment.SubjectID}
		subject, err := l.svcCtx.CommentRpc.GetSubject(l.ctx, &inGetSubject)
		if err != nil {
			l.Logger.Errorf("[RPC] Get subject info error", err)
			return nil, errors.New(http.StatusInternalServerError, "发生未知错误")
		}
		// check message type
		if subject.MessageType == biz.MessageTypePost {
			inGetPost := post.GetPostDetailRequest{PostID: subject.MessageID}
			postDetail, err := l.svcCtx.PostRpc.GetPostDetail(l.ctx, &inGetPost)
			if err != nil {
				l.Logger.Errorf("[RPC] Get post detail error", err)
				return nil, errors.New(http.StatusInternalServerError, "发生未知错误")
			}
			if uid != postDetail.AuthorID {
				return nil, errors.New(http.StatusUnauthorized, "你没有权限删除本评论")
			}

		} else if subject.MessageType == biz.MessageTypeVideo {
			// todo: video

		} else {
			return nil, errors.New(http.StatusInternalServerError, "消息类型错误")
		}
	} else {
		return nil, errors.New(http.StatusBadRequest, "role不存在")
	}

	// delete comment
	in := comment.DeleteCommentRequest{
		CommentID: req.CommentID,
		Operator:  req.Operator,
	}
	_, err = l.svcCtx.CommentRpc.DeleteComment(l.ctx, &in)
	if err != nil {
		l.Logger.Errorf("[RPC] Delete comment error", err)
	}

	return &types.DeleteCommentResponse{}, nil
}
