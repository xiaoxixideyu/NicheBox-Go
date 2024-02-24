package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/common/biz"
	"nichebox/service/comment/model/dto"
	"nichebox/service/comment/rpc/pb/comment"
	"nichebox/service/post/rpc/pb/post"
	"nichebox/service/push/rpc/pb/push"

	"nichebox/service/comment/api/internal/svc"
	"nichebox/service/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentLogic {
	return &CommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentLogic) Comment(req *types.CommentRequest) (resp *types.CommentResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	in := comment.CommentRequest{
		Uid:         uid,
		MessageID:   req.MessageID,
		MessageType: int32(req.MessageType),
		RootID:      req.RootID,
		ParentID:    req.ParentID,
		DialogID:    req.DialogID,
		Content:     req.Content,
	}
	out, err := l.svcCtx.CommentRpc.Comment(l.ctx, &in)
	if err != nil {
		l.Logger.Errorf("[RPC] Add comment error", err)
		return nil, errors.New(http.StatusInternalServerError, "发生未知错误")
	}
	resp = &types.CommentResponse{Comment: types.CommentInfo{
		CommentID:       out.Comment.CommentID,
		SubjectID:       out.Comment.SubjectID,
		RootID:          out.Comment.RootID,
		ParentID:        out.Comment.ParentID,
		DialogID:        out.Comment.DialogID,
		OwnerID:         out.Comment.DialogID,
		LikeCount:       int(out.Comment.LikeCount),
		ThumbsUp:        out.Comment.ThumbsUp,
		Floor:           int(out.Comment.Floor),
		CreateTime:      out.Comment.CreateTime,
		InnerFloorCount: int(out.Comment.InnerFloorCount),
		Content:         out.Comment.Content,
	}}

	// notification
	// notify post/video author
	var authorID int64
	var getMessageDetailSuccess = true
	var info string
	if in.MessageType == biz.MessageTypePost {
		inPost := post.GetPostDetailRequest{PostID: in.MessageID}
		out, err := l.svcCtx.PostRpc.GetPostDetail(l.ctx, &inPost)
		if err != nil {
			l.Logger.Errorf("[Rpc] Get post detail failed, err:", err)
			getMessageDetailSuccess = false
		} else {
			authorID = out.AuthorID
			info = NewCommentNotificationToPostAuthorInfo
		}

	} else if in.MessageType == biz.MessageTypeVideo {

	}
	if getMessageDetailSuccess {
		msg := dto.CommentNotificationMessage{
			NewCommentOwner: uid,
			Info:            info,
		}
		bytes, err := json.Marshal(&msg)
		if err != nil {
			l.Logger.Errorf("[Json] Marshal failed, err:", err)
		} else {
			inPush := push.PushToUserRequest{
				Uid:  authorID,
				Data: bytes,
			}
			_, err := l.svcCtx.PushRpc.PushToUser(l.ctx, &inPush)
			if err != nil {
				l.Logger.Errorf("[RPC] Notification push to user failed, err:", err)
			}
		}
	}

	// todo: notify root cmt owner and parent cmt owner

	resp.Comment.InnerFloorComments = make([]*types.CommentInfo, 0, len(out.Comment.InnerFloorComments))
	for _, cmt := range out.Comment.InnerFloorComments {
		innerCmt := types.CommentInfo{
			CommentID:       cmt.CommentID,
			SubjectID:       cmt.SubjectID,
			RootID:          cmt.RootID,
			ParentID:        cmt.ParentID,
			DialogID:        cmt.DialogID,
			OwnerID:         cmt.DialogID,
			LikeCount:       int(cmt.LikeCount),
			ThumbsUp:        cmt.ThumbsUp,
			Floor:           int(cmt.Floor),
			CreateTime:      cmt.CreateTime,
			InnerFloorCount: int(cmt.InnerFloorCount),
			Content:         cmt.Content,
		}
		resp.Comment.InnerFloorComments = append(resp.Comment.InnerFloorComments, &innerCmt)
	}

	return resp, nil
}
