package redis

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"nichebox/common/biz"
	"nichebox/service/comment/model"
	"strconv"
	"time"
)

type RedisInterface struct {
	rds *redis.Redis
}

func (r *RedisInterface) DeleteCommentCtx(ctx context.Context, commentID int64) error {
	key := KeyPrefixComment + KeyContent + strconv.FormatInt(commentID, 10)
	_, err := r.rds.DelCtx(ctx, key)
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisInterface) SetSubjectInfoCtx(ctx context.Context, subject *model.Subject) error {
	keyMsg := KeyPrefixComment + KeySubjectByMessage + strconv.FormatInt(int64(subject.TypeID), 10) + Separator + strconv.FormatInt(subject.MessageID, 10)
	keySbj := KeyPrefixComment + KeySubject + strconv.Itoa(int(subject.ID))

	subjectInfo, err := json.Marshal(subject)
	if err != nil {
		return err
	}

	err = r.rds.SetCtx(ctx, keyMsg, string(subjectInfo))
	if err != nil {
		return err
	}
	err = r.rds.SetCtx(ctx, keySbj, string(subjectInfo))
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisInterface) SetInnerFloorCommentIDs(ctx context.Context, rootID string, comments []*model.Comment) error {
	key := KeyPrefixComment + KeyInnerFloor + KeyCreateTime + rootID
	zs := make([]*redis.Z, 0, len(comments))
	for _, c := range comments {
		//z := redis.Z{
		//	Score:  float64(c.LikeCount),
		//	Member: c.CommentID,
		//}
		z := redis.Z{
			Score:  float64(c.CreatedAt.Unix()),
			Member: c.CommentID,
		}
		zs = append(zs, &z)
	}
	err := r.rds.PipelinedCtx(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.ZAdd(ctx, key, zs...)
		_, err := pipeliner.Exec(ctx)
		return err
	})
	return err
}

func (r *RedisInterface) SetCommentIndexesWithScoreBySubjectIDCtx(ctx context.Context, subjectID int64, comments []*model.Comment) error {
	keyFloor := KeyPrefixComment + KeyCommentIndex + KeyFloor + strconv.FormatInt(subjectID, 10)
	keyLikeCount := KeyPrefixComment + KeyCommentIndex + KeyLikeCount + strconv.FormatInt(subjectID, 10)
	zFloors := make([]*redis.Z, 0, len(comments))
	zLikeCounts := make([]*redis.Z, 0, len(comments))
	for _, c := range comments {
		zFloor := redis.Z{
			Score:  float64(c.Floor),
			Member: c.CommentID,
		}
		zLikeCount := redis.Z{
			Score:  float64(c.LikeCount),
			Member: c.CommentID,
		}
		zFloors = append(zFloors, &zFloor)
		zLikeCounts = append(zLikeCounts, &zLikeCount)
	}
	err := r.rds.PipelinedCtx(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.ZAdd(ctx, keyFloor, zFloors...)
		pipeliner.ZAdd(ctx, keyLikeCount, zLikeCounts...)
		_, err := pipeliner.Exec(ctx)
		return err
	})
	return err
}

func (r *RedisInterface) BatchSetCommentsCtx(ctx context.Context, caches []*model.CommentCache) error {
	err := r.rds.PipelinedCtx(ctx, func(pipeliner redis.Pipeliner) error {
		for _, cache := range caches {
			id := strconv.FormatInt(cache.CommentID, 10)
			key := KeyPrefixComment + KeyContent + id
			bytes, err := json.Marshal(cache)
			if err != nil {
				continue
			}
			pipeliner.Set(ctx, key, bytes, NormalExpiration*time.Second)
		}
		_, err := pipeliner.Exec(ctx)
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisInterface) GetInnerFloorCommentIDs(ctx context.Context, rootID string, start, stop int) ([]string, error) {
	key := KeyPrefixComment + KeyInnerFloor + KeyCreateTime + rootID
	vals, err := r.rds.ZrevrangeCtx(ctx, key, int64(start), int64(stop))
	if err != nil {
		return nil, err
	}
	if len(vals) == 0 {
		return nil, redis.Nil
	}
	return vals, nil
}

func (r *RedisInterface) BatchGetCommentsByIDsCtx(ctx context.Context, ids []string) (map[string]string, []string, error) {
	errIDs := make([]string, 0, len(ids))
	vals := make(map[string]string)

	for _, id := range ids {
		key := KeyPrefixComment + KeyContent + id
		val, err := r.rds.GetCtx(ctx, key)

		if err != nil || val == "" {
			// record the id that cache expired or occurred error
			errIDs = append(errIDs, id)
		} else {
			// record values
			vals[id] = val
		}
	}

	return vals, errIDs, nil
}

func (r *RedisInterface) GetCommentIndexesWithScoreBySubjectIDCtx(ctx context.Context, subjectID int64, page, size int, order string) ([]redis.Pair, error) {
	var err error
	var indexes []redis.Pair
	var key string

	start := (page - 1) * size
	stop := start + size - 1

	if order == biz.OrderByTimeAsc {
		key = KeyPrefixComment + KeyCommentIndex + KeyFloor + strconv.FormatInt(subjectID, 10)
		indexes, err = r.rds.ZrangeWithScoresCtx(ctx, key, int64(start), int64(stop))

	} else if order == biz.OrderByTimeDesc {
		key = KeyPrefixComment + KeyCommentIndex + KeyFloor + strconv.FormatInt(subjectID, 10)
		indexes, err = r.rds.ZrevrangeWithScoresCtx(ctx, key, int64(start), int64(stop))

	} else if order == biz.OrderByLikeCount {
		key = KeyPrefixComment + KeyCommentIndex + KeyLikeCount + strconv.FormatInt(subjectID, 10)
		indexes, err = r.rds.ZrevrangeWithScoresCtx(ctx, key, int64(start), int64(stop))

	} else {
		return nil, biz.ErrRedisUnknownOrder
	}

	if err != nil {
		return nil, err
	}
	if len(indexes) == 0 {
		// check if page/size out of size or redis cache expired
		card, err := r.rds.ZcardCtx(ctx, key)
		if err != nil {
			return nil, err
		}
		if card == 0 {
			// cache expired or no data in this subject
			return nil, redis.Nil
		} else {
			// page/size out of size
			return nil, biz.ErrRedisOutOfBounds
		}
	}
	return indexes, nil
}

func (r *RedisInterface) GetSubjectInfoByMessageCtx(ctx context.Context, messageID int64, messageType int) (string, error) {
	key := KeyPrefixComment + KeySubjectByMessage + strconv.FormatInt(int64(messageType), 10) + Separator + strconv.FormatInt(messageID, 10)
	val, err := r.rds.GetCtx(ctx, key)
	if err != nil {
		return "", err
	}
	if val == "" {
		return "", redis.Nil
	}
	return val, nil
}

func (r *RedisInterface) GetSubjectInfoBySubjectIDCtx(ctx context.Context, subjectID int64) (string, error) {
	key := KeyPrefixComment + KeySubject + strconv.FormatInt(subjectID, 10)
	val, err := r.rds.GetCtx(ctx, key)
	if err != nil {
		return "", err
	}
	if val == "" {
		return "", redis.Nil
	}
	return val, nil
}

func (r *RedisInterface) GetCommentCtx(ctx context.Context, commentID int64) (string, error) {
	key := KeyPrefixComment + KeyContent + strconv.FormatInt(commentID, 10)
	val, err := r.rds.GetCtx(ctx, key)
	if err != nil {
		return "", err
	}
	if val == "" {
		return "", redis.Nil
	}
	return val, nil
}

func (r *RedisInterface) DeleteSubjectInfoCtx(ctx context.Context, subjectID int64) error {
	keySbj := KeyPrefixComment + KeySubject + strconv.FormatInt(subjectID, 10)

	val, _ := r.rds.GetCtx(ctx, keySbj)
	subjectInfo := model.Subject{}
	json.Unmarshal([]byte(val), &subjectInfo)

	keyMsg := KeyPrefixComment + KeySubjectByMessage + strconv.FormatInt(int64(subjectInfo.TypeID), 10) + Separator + strconv.FormatInt(subjectInfo.MessageID, 10)
	_, err := r.rds.DelCtx(ctx, keySbj)
	if err != nil {
		return err
	}
	_, err = r.rds.DelCtx(ctx, keyMsg)
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisInterface) DeleteCommentsBySubjectIDCtx(ctx context.Context, subjectID int64) error {
	keyFloor := KeyPrefixComment + KeyCommentIndex + KeyFloor + strconv.FormatInt(subjectID, 10)
	keyLikeCount := KeyPrefixComment + KeyCommentIndex + KeyLikeCount + strconv.FormatInt(subjectID, 10)
	_, err := r.rds.ZremCtx(ctx, keyFloor)
	if err != nil {
		return err
	}
	_, err = r.rds.ZremCtx(ctx, keyLikeCount)
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisInterface) DeleteInnerFloorCommentsByRootIDCtx(ctx context.Context, rootID int64) error {
	key := KeyPrefixComment + KeyInnerFloor + strconv.FormatInt(rootID, 10)
	_, err := r.rds.ZremCtx(ctx, key)
	if err != nil {
		return err
	}
	return nil
}

func NewRedisInterface(hosts []string, deployType, pass string, tls, nonBlock bool, pingTimeout int) (model.CommentCacheInterface, error) {
	conf := redis.RedisConf{
		// todo: judge node or cluster
		Host:        hosts[0],
		Type:        deployType,
		Pass:        pass,
		Tls:         tls,
		NonBlock:    nonBlock,
		PingTimeout: time.Duration(pingTimeout) * time.Second,
	}
	r := &RedisInterface{
		rds: redis.MustNewRedis(conf),
	}
	return r, nil
}
