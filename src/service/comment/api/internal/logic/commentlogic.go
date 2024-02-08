package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	"net/http"
	"nichebox/service/comment/rpc/pb/comment"

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
	// todo: get root like count and sub like count

	// todo: notification

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
