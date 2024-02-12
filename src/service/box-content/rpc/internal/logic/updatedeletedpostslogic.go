package logic

import (
	"context"
	"nichebox/service/box-content/model"
	"time"

	"nichebox/service/box-content/rpc/internal/svc"
	"nichebox/service/box-content/rpc/pb/box-content"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeletedPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDeletedPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeletedPostsLogic {
	return &UpdateDeletedPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDeletedPostsLogic) UpdateDeletedPosts(in *box_content.UpdateDeletedPostsRequest) (*box_content.UpdateDeletedPostsResponse, error) {
	var err error

	for i := 0; i < len(in.DeletedPosts); {
		count := int(in.DeletedPosts[i].InfoCount)
		boxID := in.DeletedPosts[i].BoxID
		j := i
		infos := make([]*model.ModifiedPostInfo, 0, count)

		for ; j < i+count; j++ {
			deletedPost := in.DeletedPosts[j]

			t, _ := time.ParseInLocation(time.DateTime, deletedPost.Time, time.Local)
			info := model.ModifiedPostInfo{
				MessageID: deletedPost.PostID,
				Time:      t,
			}
			infos = append(infos, &info)
		}
		errCache := l.svcCtx.BoxContentCacheInterface.UpdateDeletedPostsCtx(l.ctx, boxID, infos)
		if errCache != nil {
			l.Logger.Errorf("[Redis] Update deleted post infos error", errCache)
			err = errCache
		}

		i = j
	}

	if err != nil {
		return nil, err
	}

	return &box_content.UpdateDeletedPostsResponse{}, nil
}
