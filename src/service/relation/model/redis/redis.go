package redis

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"nichebox/service/relation/model"
	"strconv"
	"time"
)

type RedisInterface struct {
	rds *redis.Redis
}

func (r RedisInterface) GetAllRelationshipsCtx(ctx context.Context, uid int64) ([]*model.CacheRelationshipAttribute, error) {
	key := KeyPrefixRelation + KeyRelationships + strconv.FormatInt(uid, 10)
	vals, err := r.rds.HvalsCtx(ctx, key)
	if err != nil {
		return nil, err
	}
	if len(vals) == 0 {
		return nil, redis.Nil
	}
	attrs := make([]*model.CacheRelationshipAttribute, 0, len(vals))
	for _, val := range vals {
		attr := model.CacheRelationshipAttribute{}
		err := json.Unmarshal([]byte(val), &attr)
		if err != nil {
			return nil, err
		}
		attrs = append(attrs, &attr)
	}
	return attrs, nil
}

func (r RedisInterface) GetRelationCountCtx(ctx context.Context, uid int64) (followers, followings int, err error) {
	key := KeyPrefixRelation + KeyCount + strconv.FormatInt(uid, 10)
	val, err := r.rds.GetCtx(ctx, key)
	if err != nil {
		return 0, 0, err
	}
	if val == "" {
		return 0, 0, redis.Nil
	}
	followers, followings = model.DecodeRelationCountCache(val)
	return followers, followings, nil
}

func (r RedisInterface) DeleteRelationCountCtx(ctx context.Context, uid int64) error {
	key := KeyPrefixRelation + KeyCount + strconv.FormatInt(uid, 10)
	_, err := r.rds.DelCtx(ctx, key)
	return err
}

func (r RedisInterface) SetRelationCountCtx(ctx context.Context, uid int64, followers, followings int) error {
	key := KeyPrefixRelation + KeyCount + strconv.FormatInt(uid, 10)
	val := model.EncodeRelationCountCache(followers, followings)
	err := r.rds.SetCtx(ctx, key, val)
	return err
}

func (r RedisInterface) GetRelationshipCtx(ctx context.Context, uid int64, fid int64) (*model.CacheRelationshipAttribute, error) {
	key := KeyPrefixRelation + KeyRelationships + strconv.FormatInt(uid, 10)
	hKey := strconv.FormatInt(fid, 10)
	val, err := r.rds.HgetCtx(ctx, key, hKey)
	if err != nil {
		return nil, err
	}
	if val == "" {
		return nil, redis.Nil
	}
	attr := model.CacheRelationshipAttribute{}
	err = json.Unmarshal([]byte(val), &attr)
	if err != nil {
		return nil, err
	}
	return &attr, nil
}

func (r RedisInterface) RemoveRelationshipsCtx(ctx context.Context, uid int64) error {
	key := KeyPrefixRelation + KeyRelationships + strconv.FormatInt(uid, 10)
	_, err := r.rds.DelCtx(ctx, key)
	if err != nil {
		return err
	}
	return nil
}

func (r RedisInterface) BatchAddRelationshipsCtx(ctx context.Context, uid int64, attrs []*model.CacheRelationshipAttribute) error {
	key := KeyPrefixRelation + KeyRelationships + strconv.FormatInt(uid, 10)
	vals := make(map[string]string, len(attrs))
	for _, attr := range attrs {
		bytes, err := json.Marshal(attr)
		if err != nil {
			return err
		}
		id := strconv.FormatInt(attr.Fid, 10)
		vals[id] = string(bytes)
	}
	err := r.rds.HmsetCtx(ctx, key, vals)
	return err
}

func NewRedisInterface(hosts []string, deployType, pass string, tls, nonBlock bool, pingTimeout int) (model.RelationCacheInterface, error) {
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
