package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nichebox/common/biz"
	"nichebox/service/comment/model"
	"nichebox/service/comment/model/dto"
	"nichebox/service/comment/rpc/internal/svc"
	"nichebox/service/comment/rpc/pb/comment"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsFromSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentsFromSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsFromSubjectLogic {
	return &GetCommentsFromSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentsFromSubjectLogic) GetCommentsFromSubject(in *comment.GetCommentsFromSubjectRequest) (*comment.GetCommentsFromSubjectResponse, error) {
	subject, err := l.querySubject(in)
	if err != nil {
		return nil, err
	}
	/*
		// 维护需要查DB的needDBAll和needDBAlone切片
		// 维护可以查cache的root切片、sub map（key=rootID,value=[]commentID）
		// 维护root和sub公用的content map、info map
		// 1. 从commentindex表下找root评论索引
		// ① 如果找到，加入到root切片，进入2.
		// ② 如果找不到，直接调用根据subject查找的dao方法（全量走数据库）
		// 2. 根据root索引找到自己的innerfloor索引
		// ① 如果找到，加入到sub map，进入3.
		// ② 如果找不到，将rootID加入到needDBAll中
		// 3. 逐一从root切片中向cache查找content、info
		// ① 如果找到，加入到content map、info map中，进入4
		// ② 如果找不到，将commentID加入到needDBAll
		// 4. 逐一从sub map中向cache查找sub的content、info
		// ① 如果找到，加入到content map、info map中
			② 如果找不到，将subID加入到needDBAlone
			5. DB批量查询needDB，查自己的commentID和3个sub的信息，加入到content map和info map，还有sub map（回写cache）
			6. DB批量查询needDBAlone，加入到content map和info map（回写cache）
			7. 遍历root切片，组装rpc comment info数组
			8. 根据subMap的对应关系组装最终的rpc comment info
	*/

	// query root comment ids from cache
	rootIDs, err := l.queryRootCommentIDsFromCache(subject, in)
	if err != nil {
		if errors.Is(err, biz.ErrRedisOutOfBounds) {
			// page/size too large
			return nil, status.Error(codes.OutOfRange, err.Error())
		}
		if !errors.Is(err, redis.Nil) {
			l.Logger.Errorf("[Redis] Get comment indexes error", err)
		}
		rootIDs, infos, contents, innerFloorIDMap, err := l.queryRootAndInnerFloorCommentsAndContentsBySubjectIDFromDB(int64(subject.ID), in)
		if err != nil {
			return nil, err
		}
		err = l.pushRebuildSubjectCommentIndexesCacheMessage(int64(subject.ID))
		err = l.pushRebuildInnerFloorCommentIndexesCacheMessage(rootIDs)
		err = l.rewriteCommentCache(infos, contents)
		// todo: handle error

		// assemble rpc result
		rpcCommentInfos := make([]*comment.CommentInfo, 0, len(rootIDs))
		for _, rootIDFromDB := range rootIDs {
			rootID := strconv.FormatInt(rootIDFromDB, 10)
			commentInfo := l.assembleRpcCommentInfoFromCacheMapOrDBMap(rootID, infos, contents, nil)
			commentInfo.InnerFloorComments = make([]*comment.CommentInfo, 0, MaxAmountsOfSubCommentsToShow)
			for _, innerID := range innerFloorIDMap[rootID] {
				subCommentInfo := l.assembleRpcCommentInfoFromCacheMapOrDBMap(innerID, infos, contents, nil)
				commentInfo.InnerFloorComments = append(commentInfo.InnerFloorComments, subCommentInfo)
			}
			rpcCommentInfos = append(rpcCommentInfos, commentInfo)
		}

		return &comment.GetCommentsFromSubjectResponse{Comments: rpcCommentInfos}, nil
	}

	needQueryDBWithSub := make(map[string]struct{})
	needQueryDBAlone := make(map[string]struct{})
	needRebuildInnerFloorIndexCacheID := make([]int64, 0, len(rootIDs))
	innerFloorIDMap := make(map[string][]string)

	infos := make(map[string]*model.Comment)
	contents := make(map[string]*model.CommentContent)
	caches := make(map[string]*model.CommentCache)

	// query inner floor comment ids from cache
	l.queryInnerFloorCommentIDsFromCache(rootIDs, innerFloorIDMap, needQueryDBWithSub, &needRebuildInnerFloorIndexCacheID)
	// push to mq to rebuild inner floor comment index
	l.pushRebuildInnerFloorCommentIndexesCacheMessage(needRebuildInnerFloorIndexCacheID)

	// query root comment from cache
	l.queryRootCommentsAndContentsFromCache(rootIDs, needQueryDBWithSub, caches)

	// query inner floor comment from cache
	l.queryInnerFloorCommentsAndContentsFromCache(innerFloorIDMap, needQueryDBAlone, caches)

	// query comment and its inner floor comment infos and contents from db
	err = l.queryRootWithSubCommentsAndContentsFromDB(needQueryDBWithSub, infos, contents, innerFloorIDMap)
	if err != nil {
		return nil, err
	}

	// query single comment info and content from db
	err = l.queryCommentsAndContentsFromDB(needQueryDBAlone, infos, contents)
	if err != nil {
		return nil, err
	}

	go func() {
		l.rewriteCommentCache(infos, contents)
	}()

	// assemble rpc result
	rpcCommentInfos := make([]*comment.CommentInfo, 0, len(rootIDs))
	for _, rootID := range rootIDs {
		commentInfo := l.assembleRpcCommentInfoFromCacheMapOrDBMap(rootID, infos, contents, caches)
		commentInfo.InnerFloorComments = make([]*comment.CommentInfo, 0, MaxAmountsOfSubCommentsToShow)
		for _, innerID := range innerFloorIDMap[rootID] {
			subCommentInfo := l.assembleRpcCommentInfoFromCacheMapOrDBMap(innerID, infos, contents, caches)
			commentInfo.InnerFloorComments = append(commentInfo.InnerFloorComments, subCommentInfo)
		}
		rpcCommentInfos = append(rpcCommentInfos, commentInfo)
	}

	return &comment.GetCommentsFromSubjectResponse{Comments: rpcCommentInfos}, nil
}

func (l *GetCommentsFromSubjectLogic) queryInnerFloorCommentsAndContentsFromCache(mInnerFloor map[string][]string, s map[string]struct{}, mCache map[string]*model.CommentCache) {
	for _, subIDs := range mInnerFloor {
		for i, subID := range subIDs {
			if i == MaxAmountsOfSubCommentsToShow {
				break
			}
			subIDInt64, _ := strconv.ParseInt(subID, 10, 64)
			val, err := l.svcCtx.CommentCacheInterface.GetCommentCtx(l.ctx, subIDInt64)
			if err != nil {
				s[subID] = struct{}{}
			} else {
				var cache model.CommentCache
				err := json.Unmarshal([]byte(val), &cache)
				if err != nil {
					l.Logger.Errorf("[Json] Unmarshal sub comment cache error", err)
					s[subID] = struct{}{}
				} else {
					// add to cache map
					mCache[subID] = &cache
				}
			}
		}
	}
}

func (l *GetCommentsFromSubjectLogic) queryInnerFloorCommentIDsFromCache(rootIDs []string, m map[string][]string, s map[string]struct{}, sr *[]int64) {
	for _, id := range rootIDs {
		innerFloorIDs, err := l.svcCtx.CommentCacheInterface.GetInnerFloorCommentIDs(l.ctx, id, 0, MaxAmountsOfSubCommentsToShow-1)
		if err != nil {
			s[id] = struct{}{}
			idInt64, _ := strconv.ParseInt(id, 10, 64)
			*sr = append(*sr, idInt64)
		} else {
			m[id] = innerFloorIDs
		}
	}

}

func (l *GetCommentsFromSubjectLogic) queryRootCommentsAndContentsFromCache(ids []string, s map[string]struct{}, m map[string]*model.CommentCache) {
	vals, errRootIDs, err := l.svcCtx.CommentCacheInterface.BatchGetCommentsByIDsCtx(l.ctx, ids)
	if err != nil {
		// copy all ids to errRootIDs
		errRootIDs = make([]string, 0, len(ids))
		copy(errRootIDs, ids)

	} else {
		// try to use comment cache
		for id, val := range vals {
			var rootCommentCache model.CommentCache
			err := json.Unmarshal([]byte(val), &rootCommentCache)
			if err != nil {
				l.Logger.Errorf("[Json] Unmarshal comment info cache error", err)
				// add to errRootIDs
				errRootIDs = append(errRootIDs, id)
			} else {
				// add to cache map
				m[id] = &rootCommentCache
			}
		}
	}

	for _, id := range errRootIDs {
		s[id] = struct{}{}
	}
}

func (l *GetCommentsFromSubjectLogic) querySubject(in *comment.GetCommentsFromSubjectRequest) (*model.Subject, error) {
	// query subject
	subjectStr, err := l.svcCtx.CommentCacheInterface.GetSubjectInfoByMessageCtx(l.ctx, in.MessageID, int(in.MessageType))
	var subject model.Subject
	needQueryDBSubject := false
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			l.Logger.Errorf("[Redis] Get subject info by message error", err)
		}
		needQueryDBSubject = true
	}
	if !needQueryDBSubject {
		// try to use subject cache
		err := json.Unmarshal([]byte(subjectStr), &subject)
		if err != nil {
			// cache data invalid, delete cache and query db
			l.Logger.Errorf("[Json] Unmarshal subject info cache error", err)
			needQueryDBSubject = true
		}
	}

	if needQueryDBSubject {
		subject = model.Subject{
			MessageID: in.MessageID,
			TypeID:    int(in.MessageType),
		}
		err := l.svcCtx.CommentInterface.FirstOrCreateSubject(&subject)
		if err != nil {
			l.Logger.Errorf("[MySQL] Get subject error", err)
			return nil, err
		}
		// rewrite subject cache
		l.svcCtx.CommentCacheInterface.SetSubjectInfoCtx(l.ctx, &subject)
	}
	return &subject, nil
}

func (l *GetCommentsFromSubjectLogic) queryRootCommentIDsFromCache(subject *model.Subject, in *comment.GetCommentsFromSubjectRequest) ([]string, error) {
	kvs, err := l.svcCtx.CommentCacheInterface.GetCommentIndexesWithScoreBySubjectIDCtx(l.ctx, int64(subject.ID), int(in.Page), int(in.Size), in.Order)
	if err != nil {
		return nil, err
	}
	ids := make([]string, 0, len(kvs))
	for _, kv := range kvs {
		ids = append(ids, kv.Key)
	}
	return ids, nil
}

func (l *GetCommentsFromSubjectLogic) queryRootWithSubCommentsAndContentsFromDB(s map[string]struct{}, infos map[string]*model.Comment, contents map[string]*model.CommentContent, innerFloorIDMap map[string][]string) error {
	rootIDs := make([]int64, 0, len(s))
	for k, _ := range s {
		rootID, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			l.Logger.Errorf("[Strconv] Convert string to int64 error", err)
			return err
		}
		rootIDs = append(rootIDs, rootID)
	}

	rootInfos, err := l.svcCtx.CommentInterface.BatchGetComments(rootIDs)
	if err != nil {
		l.Logger.Errorf("[MySQL] Get root comment infos error", err)
		return err
	}
	rootContents, err := l.svcCtx.CommentInterface.BatchGetCommentsContents(rootIDs)
	if err != nil {
		l.Logger.Errorf("[MySQL] Get root comment contents error", err)
		return err
	}

	for i := 0; i < len(rootIDs); i++ {
		innerInfos, innerContents, err := l.svcCtx.CommentInterface.GetInnerFloorCommentsAndContentsByRootID(rootIDs[i], 1, MaxAmountsOfSubCommentsToShow)
		if err != nil {
			l.Logger.Errorf("[MySQL] Get inner floor comment infos and contents error", err)
			return err
		}
		rootIDStr := strconv.FormatInt(rootIDs[i], 10)
		infos[rootIDStr] = rootInfos[i]
		contents[rootIDStr] = rootContents[i]
		innerFloorIDMap[rootIDStr] = make([]string, 0, MaxAmountsOfSubCommentsToShow)
		for j := 0; j < len(innerInfos); j++ {
			subCommentID := strconv.FormatInt(innerInfos[j].CommentID, 10)
			infos[subCommentID] = innerInfos[j]
			contents[subCommentID] = innerContents[j]
			innerFloorIDMap[rootIDStr] = append(innerFloorIDMap[rootIDStr], subCommentID)
		}
	}

	return nil
}

func (l *GetCommentsFromSubjectLogic) queryCommentsAndContentsFromDB(s map[string]struct{}, mInfos map[string]*model.Comment, mContents map[string]*model.CommentContent) error {
	ids := make([]int64, 0, len(s))
	for k, _ := range s {
		id, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			l.Logger.Errorf("[Strconv] Convert string to int64 error", err)
			return err
		}
		ids = append(ids, id)
	}

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

func (l *GetCommentsFromSubjectLogic) rewriteCommentCache(mInfos map[string]*model.Comment, mContents map[string]*model.CommentContent) error {
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

func (l *GetCommentsFromSubjectLogic) pushRebuildSubjectCommentIndexesCacheMessage(subjectID int64) error {
	msg := dto.RebuildCacheSubjectCommentIndexMessage{
		CreateDate: time.Now(),
		SubjectID:  subjectID,
	}
	bytes, err := json.Marshal(&msg)
	if err != nil {
		l.Logger.Errorf("[Json][Producer] Json marshal error", err)
		return err
	}
	err = l.svcCtx.KqRebuildSubjectCommentIndexPusherClient.Push(string(bytes))
	if err != nil {
		l.Logger.Errorf("[Kafka][Producer] MQ push error", err)
		return err
	}
	return nil
}

func (l *GetCommentsFromSubjectLogic) pushRebuildInnerFloorCommentIndexesCacheMessage(rootIDs []int64) error {
	msg := dto.RebuildCacheInnerFloorCommentIndexMessage{RootIDs: rootIDs}
	bytes, err := json.Marshal(&msg)
	if err != nil {
		l.Logger.Errorf("[Json][Producer] Json marshal error", err)
		return err
	}
	err = l.svcCtx.KqRebuildCacheInnerFloorCommentIndexPusherClient.Push(string(bytes))
	if err != nil {
		l.Logger.Errorf("[Kafka][Producer] MQ push error", err)
		return err
	}
	return nil
}

func (l *GetCommentsFromSubjectLogic) queryRootAndInnerFloorCommentsAndContentsBySubjectIDFromDB(subjectID int64, in *comment.GetCommentsFromSubjectRequest) ([]int64, map[string]*model.Comment, map[string]*model.CommentContent, map[string][]string, error) {
	rootComments, err := l.svcCtx.CommentInterface.GetRootCommentsBySubjectID(subjectID, int(in.Page), int(in.Size), in.Order)
	if err != nil {
		l.Logger.Errorf("[MySQL] Get root comments by subject id from db error", err)
		return nil, nil, nil, nil, err
	}
	rootIDs := make([]int64, 0, len(rootComments))
	for _, c := range rootComments {
		rootIDs = append(rootIDs, c.CommentID)
	}
	subComments, err := l.svcCtx.CommentInterface.BatchGetInnerFloorComments(rootIDs, 1, MaxAmountsOfSubCommentsToShow)
	if err != nil {
		l.Logger.Errorf("[MySQL] Get inner floor comments by root ids from db error", err)
		return nil, nil, nil, nil, err
	}
	m := make(map[string][]string)
	allIDs := make([]int64, 0, len(subComments)+len(rootComments))
	allIDs = append(allIDs, rootIDs...)
	for _, sub := range subComments {
		subRootID := strconv.FormatInt(sub.RootID, 10)
		_, ok := m[subRootID]
		if !ok {
			m[subRootID] = make([]string, 0, MaxAmountsOfSubCommentsToShow)
		}
		m[subRootID] = append(m[subRootID], strconv.FormatInt(sub.CommentID, 10))
		allIDs = append(allIDs, sub.CommentID)
	}
	comments := make([]*model.Comment, 0, len(subComments)+len(rootComments))
	comments = append(comments, rootComments...)
	comments = append(comments, subComments...)

	commentContents, err := l.svcCtx.CommentInterface.BatchGetCommentsContents(allIDs)
	if err != nil {
		l.Logger.Errorf("[MySQL] Get comment contents error", err)
		return nil, nil, nil, nil, err
	}

	infos := make(map[string]*model.Comment)
	contents := make(map[string]*model.CommentContent)

	for i, id := range allIDs {
		idStr := strconv.FormatInt(id, 10)
		infos[idStr] = comments[i]
		contents[idStr] = commentContents[i]
	}

	return rootIDs, infos, contents, m, nil
}

func (l *GetCommentsFromSubjectLogic) assembleRpcCommentInfoFromCacheMapOrDBMap(commentID string, infos map[string]*model.Comment, contents map[string]*model.CommentContent, caches map[string]*model.CommentCache) *comment.CommentInfo {
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
			ThumbsUp:           false,
			Floor:              int32(cache.Floor),
			CreateTime:         cache.CreatedAt.Format(time.DateTime),
			InnerFloorCount:    int32(cache.InnerFloorCount),
			InnerFloorComments: nil,
			Content:            cache.Content,
		}
	} else {
		info, ok := infos[commentID]
		if !ok {
			l.Logger.Errorf("[Logic] Not found comment", commentID)
			return &commentInfo
		}
		content, ok := contents[commentID]
		if !ok {
			l.Logger.Errorf("[Logic] Not found comment", commentID)
			return &commentInfo
		}
		commentInfo = comment.CommentInfo{
			CommentID:          info.CommentID,
			SubjectID:          info.SubjectID,
			RootID:             info.RootID,
			ParentID:           info.ParentID,
			DialogID:           info.DialogID,
			OwnerID:            info.OwnerID,
			ThumbsUp:           false,
			Floor:              int32(info.Floor),
			CreateTime:         info.CreatedAt.Format(time.DateTime),
			InnerFloorCount:    int32(info.InnerFloorCount),
			InnerFloorComments: nil,
			Content:            content.Content,
		}
	}

	return &commentInfo
}
