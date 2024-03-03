package redis

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"nichebox/service/relation/model"
	"strconv"
	"time"
)

type RedisInterface struct {
	rds  *redis.Redis
	bits uint
}

func (r RedisInterface) BloomAddRelationCtx(ctx context.Context, uid int64, fid int64) error {
	// both 2 users update their own bloom filer
	bloomKey := KeyPrefixRelation + KeyPrefixBloomFilter + strconv.FormatInt(uid, 10)
	f := bloom.New(r.rds, bloomKey, r.bits)
	err := f.AddCtx(ctx, []byte(strconv.FormatInt(fid, 10)))
	if err != nil {
		return err
	}

	bloomKey = KeyPrefixRelation + KeyPrefixBloomFilter + strconv.FormatInt(fid, 10)
	f = bloom.New(r.rds, bloomKey, r.bits)
	err = f.AddCtx(ctx, []byte(strconv.FormatInt(uid, 10)))

	return err
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
	bloomKey := KeyPrefixRelation + KeyPrefixBloomFilter + strconv.FormatInt(uid, 10)
	f := bloom.New(r.rds, bloomKey, r.bits)
	exists, _ := f.ExistsCtx(ctx, []byte(strconv.FormatInt(fid, 10)))
	if !exists {
		// no relationship
		return &model.CacheRelationshipAttribute{
			Fid:          fid,
			Relationship: model.ConvertRelationNumberToString(model.RelationNone),
			UpdateTime:   time.Time{},
		}, nil
	}

	attr, err := r.getRelationshipCtx(ctx, uid, fid)
	if err != nil {
		return nil, err
	}
	if attr.Relationship == model.ConvertRelationNumberToString(model.RelationNone) {
		// hash cache stores followings and friends only, so we need to check fid's hash cache to know if it is the follower relation
		return r.getRelationshipCtx(ctx, fid, uid)
	}
	return attr, nil
}

func (r RedisInterface) getRelationshipCtx(ctx context.Context, uid, fid int64) (*model.CacheRelationshipAttribute, error) {
	key := KeyPrefixRelation + KeyRelationships + strconv.FormatInt(uid, 10)
	hKey := strconv.FormatInt(fid, 10)
	val, err := r.rds.HgetCtx(ctx, key, hKey)
	if err != nil {
		return nil, err
	}
	if val == "" {
		// check whether no relationship or hash expired
		exists, err := r.rds.ExistsCtx(ctx, key)
		if err != nil {
			return nil, err
		}
		if !exists {
			// hash expired
			return nil, redis.Nil
		} else {
			// no relationship
			return &model.CacheRelationshipAttribute{
				Fid:          fid,
				Relationship: model.ConvertRelationNumberToString(model.RelationNone),
				UpdateTime:   time.Time{},
			}, nil
		}
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

func NewRedisInterface(hosts []string, deployType, pass string, tls, nonBlock bool, pingTimeout int, bloomFilterBits uint) (model.RelationCacheInterface, error) {
	conf := redis.RedisConf{
		// todo: judge node or cluster
		Host:        hosts[0],
		Type:        deployType,
		Pass:        pass,
		Tls:         tls,
		NonBlock:    nonBlock,
		PingTimeout: time.Duration(pingTimeout) * time.Second,
	}
	rds := redis.MustNewRedis(conf)
	r := &RedisInterface{
		rds:  rds,
		bits: bloomFilterBits,
	}
	return r, nil
}
