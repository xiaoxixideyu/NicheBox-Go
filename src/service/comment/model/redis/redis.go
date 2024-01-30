package redis

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"nichebox/service/comment/model"
	"strconv"
	"time"
)

type RedisInterface struct {
	rds *redis.Redis
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
	key := KeyPrefixComment + KeySubject + strconv.FormatInt(subjectID, 10)
	_, err := r.rds.DelCtx(ctx, key)
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
