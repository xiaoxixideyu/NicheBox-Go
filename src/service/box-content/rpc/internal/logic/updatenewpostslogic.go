package logic

import (
	"context"
	"nichebox/service/box-content/model"
	"nichebox/service/box-content/rpc/internal/svc"
	"nichebox/service/box-content/rpc/pb/box-content"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNewPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNewPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNewPostsLogic {
	return &UpdateNewPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateNewPostsLogic) UpdateNewPosts(in *box_content.UpdateNewPostsRequest) (*box_content.UpdateNewPostsResponse, error) {
	var err error
	for i := 0; i < len(in.NewPosts); {
		count := int(in.NewPosts[i].InfoCount)
		boxID := in.NewPosts[i].BoxID
		j := i
		infos := make([]*model.ModifiedPostInfo, 0, count)

		for ; j < i+count; j++ {
			newPost := in.NewPosts[j]

			t, _ := time.ParseInLocation(time.DateTime, newPost.Time, time.Local)
			info := model.ModifiedPostInfo{
				MessageID: newPost.PostID,
				Time:      t,
			}
			infos = append(infos, &info)
		}
		errCache := l.svcCtx.BoxContentCacheInterface.UpdateNewPostsCtx(l.ctx, boxID, infos)
		if errCache != nil {
			l.Logger.Errorf("[Redis] Update new post infos error", errCache)
			err = errCache
		}

		i = j
	}

	if err != nil {
		return nil, err
	}

	return &box_content.UpdateNewPostsResponse{}, nil
}
