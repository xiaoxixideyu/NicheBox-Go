package redis

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"nichebox/service/like/model"
	"strconv"
	"time"
)

type RedisInterface struct {
	rds *redis.Redis
}

func (r *RedisInterface) BatchAddThumbsUpHistoryCtx(ctx context.Context, likes []*model.Like) error {
	if len(likes) == 0 {
		return nil
	}

	key := KeyPrefixLike + KeyHistory + strconv.FormatInt(likes[0].Uid, 10) + Separator + string(likes[0].TypeID)
	err := r.rds.PipelinedCtx(ctx, func(pipeliner redis.Pipeliner) error {
		data := make([]*redis.Z, 0)
		for _, l := range likes {
			d := redis.Z{
				Score:  float64(l.UpdatedAt.Unix()),
				Member: strconv.FormatInt(l.MessageID, 10),
			}
			data = append(data, &d)
		}
		pipeliner.ZAdd(ctx, key, data...)
		_, err := pipeliner.Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (r *RedisInterface) ClearAllThumbsUpHistoryCtx(ctx context.Context, typeID int, uid int64) error {
	key := KeyPrefixLike + KeyHistory + strconv.FormatInt(uid, 10) + Separator + strconv.Itoa(typeID)
	_, err := r.rds.ZremrangebyrankCtx(ctx, key, 0, -1)
	return err
}

func (r *RedisInterface) GetThumbsUpHistoryCtx(ctx context.Context, typeID int, uid int64, start int, stop int) ([]string, error) {
	key := KeyPrefixLike + KeyHistory + strconv.FormatInt(uid, 10) + Separator + strconv.Itoa(typeID)
	messageIDs, err := r.rds.ZrevrangeCtx(ctx, key, int64(start), int64(stop))
	if err != nil {
		return nil, err
	}
	return messageIDs, nil
}

func (r *RedisInterface) RemoveThumbsUpHistoryCtx(ctx context.Context, messageID int64, typeID int, uid int64) error {
	key := KeyPrefixLike + KeyHistory + strconv.FormatInt(uid, 10) + Separator + strconv.Itoa(typeID)
	_, err := r.rds.ZremCtx(ctx, key, strconv.FormatInt(messageID, 10))
	if err != nil {
		return err
	}

	return nil
}

func NewRedisInterface(hosts []string, deployType, pass string, tls, nonBlock bool, pingTimeout int) (model.LikeCacheInterface, error) {
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

func (r *RedisInterface) SetThumbsUpCountCtx(ctx context.Context, messageID int64, typeID int, count int) error {
	key := KeyPrefixLike + KeyCount + strconv.Itoa(typeID) + Separator + strconv.FormatInt(messageID, 10)
	err := r.rds.SetexCtx(ctx, key, strconv.Itoa(count), LikeCountExpiration)
	return err
}

func (r *RedisInterface) GetThumbsUpCountCtx(ctx context.Context, messageID int64, typeID int) (string, error) {
	key := KeyPrefixLike + KeyCount + strconv.Itoa(typeID) + Separator + strconv.FormatInt(messageID, 10)
	val, err := r.rds.GetCtx(ctx, key)
	if err != nil {
		return "", err
	}
	if val == "" {
		return "", redis.Nil
	}
	return val, nil
}

func (r *RedisInterface) DeleteThumbsUpCountCtx(ctx context.Context, messageID int64, typeID int) (int, error) {
	key := KeyPrefixLike + KeyCount + strconv.Itoa(typeID) + Separator + strconv.FormatInt(messageID, 10)
	val, err := r.rds.DelCtx(ctx, key)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (r *RedisInterface) AddThumbsUpHistoryCtx(ctx context.Context, messageID int64, typeID int, uid int64) error {
	key := KeyPrefixLike + KeyHistory + strconv.FormatInt(uid, 10) + Separator + strconv.Itoa(typeID)
	now := time.Now().Unix()
	_, err := r.rds.ZaddCtx(ctx, key, now, strconv.FormatInt(messageID, 10))
	if err != nil {
		return err
	}
	// remain limited length of thumbs up history Zset
	card, _ := r.rds.ZcardCtx(ctx, key)
	if card > HistoryListLength {
		r.rds.ZremrangebyrankCtx(ctx, key, HistoryListLength+1, -1)
	}
	return nil
}

//func (r *RedisInterface) IncreaseThumbsUpCountCtx(ctx context.Context, messageID int64, typeID int) (int64, error) {
//	key := KeyPrefixLike + KeyCount + strconv.Itoa(typeID) + Separator + strconv.FormatInt(messageID, 10)
//	val, err := r.rds.IncrCtx(ctx, key)
//	if err != nil {
//		return 0, err
//	}
//	return val, nil
//}
