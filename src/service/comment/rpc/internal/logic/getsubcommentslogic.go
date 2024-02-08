package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nichebox/service/comment/model"
	"strconv"
	"time"

	"nichebox/service/comment/rpc/internal/svc"
	"nichebox/service/comment/rpc/pb/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSubCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubCommentsLogic {
	return &GetSubCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSubCommentsLogic) GetSubComments(in *comment.GetSubCommentsRequest) (*comment.GetSubCommentsResponse, error) {
	innerFloorIDs, err := l.svcCtx.CommentCacheInterface.GetInnerFloorCommentIDs(l.ctx, strconv.FormatInt(in.RootID, 10), 0, -1)
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			l.Logger.Errorf("[Redis] Get inner floor comment ids error", err)
		}
		infos, contents, err := l.queryInnerFloorCommentsAndContentsFromDB(in)
		if err != nil {
			return nil, err
		}
		// rewrite cache
		l.rewriteCommentCache(infos, contents)
		// rebuild inner floor comment index cache
		l.rebuildInnerFloorCommentIndexCache(in.RootID, infos)
		// assemble rpc result
		commentInfos := l.assembleRpcCommentInfoFromCache(infos, contents)
		return &comment.GetSubCommentsResponse{SubComments: commentInfos}, nil
	}

	// paging
	startIndex := (in.Page - 1) * in.Size
	stopIndex := startIndex + in.Size
	if startIndex > int32(len(innerFloorIDs)) {
		return nil, status.Error(codes.OutOfRange, "page/size out of bounds")
	}
	if stopIndex > int32(len(innerFloorIDs)) {
		stopIndex = int32(len(innerFloorIDs))
	}
	innerFloorIDs = innerFloorIDs[startIndex:stopIndex]

	infos := make(map[string]*model.Comment)
	contents := make(map[string]*model.CommentContent)
	caches := make(map[string]*model.CommentCache)

	// query comments and contents from cache
	needQueryDBAloneIDs := l.queryCommentsAndContentsFromCache(innerFloorIDs, caches)

	// query comments and contents from db
	err = l.queryCommentsAndContentsFromDB(needQueryDBAloneIDs, infos, contents)
	if err != nil {
		return nil, err
	}

	go func() {
		l.rewriteCommentCacheByMap(infos, contents)
	}()

	// assemble rpc result
	subComments := l.assembleRpcCommentInfoFromCacheMapOrDBMap(innerFloorIDs, infos, contents, caches)

	return &comment.GetSubCommentsResponse{SubComments: subComments}, nil
}

func (l *GetSubCommentsLogic) assembleRpcCommentInfoFromCacheMapOrDBMap(innerFloorIDs []string, infos map[string]*model.Comment, contents map[string]*model.CommentContent, caches map[string]*model.CommentCache) []*comment.CommentInfo {
	commentInfos := make([]*comment.CommentInfo, 0, len(infos)+len(caches))

	for _, commentID := range innerFloorIDs {
		var commentInfo comment.CommentInfo

		cache, ok := caches[commentID]
		if ok {
			commentInfo = comment.CommentInfo{
				CommentID:          cache.CommentID,
				SubjectID:          cache.SubjectID,
				RootID:             cache.RootID,
				ParentID:           cache.ParentID,
				DialogID:           cache.DialogID,
				OwnerID:            cache.OwnerID,
				LikeCount:          int32(cache.LikeCount),
				ThumbsUp:           false,
				Floor:              int32(cache.Floor),
				CreateTime:         cache.CreatedAt.Format(time.DateTime),
				InnerFloorCount:    int32(cache.InnerFloorCount),
				InnerFloorComments: nil,
				Content:            cache.Content,
			}
		} else {
			info := infos[commentID]
			content := contents[commentID]
			commentInfo = comment.CommentInfo{
				CommentID:          info.CommentID,
				SubjectID:          info.SubjectID,
				RootID:             info.RootID,
				ParentID:           info.ParentID,
				DialogID:           info.DialogID,
				OwnerID:            info.OwnerID,
				LikeCount:          int32(info.LikeCount),
				ThumbsUp:           false,
				Floor:              int32(info.Floor),
				CreateTime:         info.CreatedAt.Format(time.DateTime),
				InnerFloorCount:    int32(info.InnerFloorCount),
				InnerFloorComments: nil,
				Content:            content.Content,
			}
		}

		commentInfos = append(commentInfos, &commentInfo)
	}

	return commentInfos
}

func (l *GetSubCommentsLogic) assembleRpcCommentInfoFromCache(infos []*model.Comment, contents []*model.CommentContent) []*comment.CommentInfo {
	commentInfos := make([]*comment.CommentInfo, 0, len(infos))

	for i := 0; i < len(infos); i++ {
		info := infos[i]
		content := contents[i]

		commentInfo := comment.CommentInfo{
			CommentID:          info.CommentID,
			SubjectID:          info.SubjectID,
			RootID:             info.RootID,
			ParentID:           info.ParentID,
			DialogID:           info.DialogID,
			OwnerID:            info.OwnerID,
			LikeCount:          int32(info.LikeCount),
			ThumbsUp:           false,
			Floor:              int32(info.Floor),
			CreateTime:         info.CreatedAt.Format(time.DateTime),
			InnerFloorCount:    int32(info.InnerFloorCount),
			InnerFloorComments: nil,
			Content:            content.Content,
		}
		commentInfos = append(commentInfos, &commentInfo)
	}

	return commentInfos
}

func (l *GetSubCommentsLogic) rebuildInnerFloorCommentIndexCache(rootID int64, subComments []*model.Comment) error {
	err := l.svcCtx.CommentCacheInterface.SetInnerFloorCommentIDs(l.ctx, strconv.FormatInt(rootID, 10), subComments)
	if err != nil {
		l.Logger.Errorf("[Redis]Set inner floor comment index error", err)
		return err
	}
	return nil
}

func (l *GetSubCommentsLogic) rewriteCommentCache(infos []*model.Comment, contents []*model.CommentContent) []*model.CommentCache {
	caches := make([]*model.CommentCache, 0, len(infos))
	for i := 0; i < len(caches); i++ {
		info := infos[i]
		content := contents[i]
		cache := model.CommentCache{
			CommentID:       info.CommentID,
			SubjectID:       info.SubjectID,
			RootID:          info.RootID,
			ParentID:        info.ParentID,
			DialogID:        info.DialogID,
			OwnerID:         info.OwnerID,
			Floor:           info.Floor,
			LikeCount:       info.LikeCount,
			InnerFloorCount: info.InnerFloorCount,
			Status:          info.Status,
			Content:         content.Content,
			CreatedAt:       info.CreatedAt,
			UpdatedAt:       info.UpdatedAt,
		}
		caches = append(caches, &cache)
	}
	return caches
}

func (l *GetSubCommentsLogic) queryCommentsAndContentsFromDB(ids []int64, mInfos map[string]*model.Comment, mContents map[string]*model.CommentContent) error {
	infos, err := l.svcCtx.CommentInterface.BatchGetComments(ids)
	if err != nil {
		l.Logger.Errorf("[MySQL] Get comment infos error", err)
		return err
	}
	contents, err := l.svcCtx.CommentInterface.BatchGetCommentsContents(ids)
	if err != nil {
		l.Logger.Errorf("[MySQL] Get comment contents error", err)
		return err
	}

	for i := 0; i < len(ids); i++ {
		idStr := strconv.FormatInt(ids[i], 10)
		mInfos[idStr] = infos[i]
		mContents[idStr] = contents[i]
	}

	return nil
}

func (l *GetSubCommentsLogic) queryCommentsAndContentsFromCache(ids []string, m map[string]*model.CommentCache) []int64 {
	vals, errIDs, err := l.svcCtx.CommentCacheInterface.BatchGetCommentsByIDsCtx(l.ctx, ids)
	if err != nil {
		l.Logger.Errorf("[Redis] Batch get comments by ids error", err)

		// copy all ids to errRootIDs
		errIDsInt64 := make([]int64, 0, len(ids))
		for _, id := range ids {
			idInt64, _ := strconv.ParseInt(id, 10, 64)
			errIDsInt64 = append(errIDsInt64, idInt64)
		}
		return errIDsInt64
	}
	// try to use comment cache
	for id, val := range vals {
		var rootCommentCache model.CommentCache
		err := json.Unmarshal([]byte(val), &rootCommentCache)
		if err != nil {
			l.Logger.Errorf("[Json] Unmarshal comment info cache error", err)
			// add to errRootIDs
			errIDs = append(errIDs, id)
		} else {
			// add to cache map
			m[id] = &rootCommentCache
		}
	}

	errIDsInt64 := make([]int64, 0, len(errIDs))
	for _, id := range errIDs {
		idInt64, _ := strconv.ParseInt(id, 10, 64)
		errIDsInt64 = append(errIDsInt64, idInt64)
	}
	return errIDsInt64
}

func (l *GetSubCommentsLogic) rewriteCommentCacheByMap(mInfos map[string]*model.Comment, mContents map[string]*model.CommentContent) error {
	caches := make([]*model.CommentCache, 0, len(mInfos))
	for k, info := range mInfos {
		content := mContents[k].Content
		cache := model.CommentCache{
			CommentID:       info.CommentID,
			SubjectID:       info.SubjectID,
			RootID:          info.RootID,
			ParentID:        info.ParentID,
			DialogID:        info.DialogID,
			OwnerID:         info.OwnerID,
			Floor:           info.Floor,
			LikeCount:       info.LikeCount,
			InnerFloorCount: info.InnerFloorCount,
			Status:          info.Status,
			Content:         content,
			CreatedAt:       info.CreatedAt,
			UpdatedAt:       info.UpdatedAt,
		}
		caches = append(caches, &cache)
	}
	err := l.svcCtx.CommentCacheInterface.BatchSetCommentsCtx(context.Background(), caches)
	if err != nil {
		l.Logger.Errorf("[Redis] Batch set comments error", err)
		return err
	}
	return nil
}

func (l *GetSubCommentsLogic) queryInnerFloorCommentsAndContentsFromDB(in *comment.GetSubCommentsRequest) ([]*model.Comment, []*model.CommentContent, error) {
	comments, contents, err := l.svcCtx.CommentInterface.GetInnerFloorCommentsAndContentsByRootID(in.RootID, int(in.Page), int(in.Size))
	if err != nil {
		l.Logger.Errorf("[MySQL] Get inner floor comments and contents by root id error", err)
		return nil, nil, err
	}

	return comments, contents, nil
}
