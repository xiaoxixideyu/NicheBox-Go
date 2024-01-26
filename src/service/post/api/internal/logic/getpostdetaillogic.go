package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"nichebox/service/post/api/internal/svc"
	"nichebox/service/post/api/internal/types"
	"nichebox/service/post/rpc/pb/post"
	"nichebox/service/user/rpc/pb/user"
	"time"
)

type GetPostDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostDetailLogic {
	return &GetPostDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostDetailLogic) GetPostDetail(req *types.GetPostDetailRequest) (resp *types.GetPostDetailResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "uid无效")
	}

	inPost := post.GetPostDetailRequest{
		PostID: req.PostID,
	}

	outPost, err := l.svcCtx.PostRpc.GetPostDetail(l.ctx, &inPost)
	if err != nil {
		rpcStatus, ok := status.FromError(err)
		if ok {
			if rpcStatus.Code() == codes.NotFound {
				return nil, errors.New(http.StatusBadRequest, "该帖子不存在")
			}
		}
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}

	resp = &types.GetPostDetailResponse{
		AuthorID:   outPost.AuthorID,
		AuthorName: "Unknown",
		BoxID:      outPost.BoxID,
		BoxName:    "Unknown",
		Title:      outPost.Title,
		Content:    outPost.Content,
		Photos:     outPost.Photos,
		UserView:   int(outPost.UserView),
	}

	// todo: rpc协程结果收集
	ddl, _ := l.ctx.Deadline()
	newDdl := ddl.Add(-10 * time.Millisecond)
	newDdlCtx, cancelFunc := context.WithDeadline(l.ctx, newDdl)
	defer cancelFunc()

	// get author user info
	go func() {
		inUser := user.GetUserBaseInfoRequest{
			Uid: outPost.AuthorID,
		}
		outUser, err := l.svcCtx.UserRpc.GetUserBaseInfo(l.ctx, &inUser)
		if err != nil {
			return
		}
		resp.AuthorName = outUser.Username
	}()

	// todo: get box info

	// increase uv
	go func() {
		inIncrUV := post.IncreaseUserViewRequest{
			PostID:    req.PostID,
			VisitorID: uid,
		}
		l.svcCtx.PostRpc.IncreaseUserView(context.Background(), &inIncrUV)
	}()

	select {
	case <-newDdlCtx.Done():

	}

	return resp, nil
}
