package model

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

const (
	RelationFollow = 1
	RelationFriend = 2
	RelationNone   = 3
)

func ConvertRelationNumberToString(r int8) string {
	switch r {
	case RelationFollow:
		return "follow"
	case RelationFriend:
		return "friend"
	case RelationNone:
		return "None"
	default:
		return "Unknown"
	}
}

func DecodeRelationCountCache(cache string) (followers, followings int) {
	split := strings.Split(cache, ",")
	i, _ := strconv.ParseInt(split[0], 10, 64)
	followers = int(i)
	i, _ = strconv.ParseInt(split[1], 10, 64)
	followings = int(i)
	return
}

func EncodeRelationCountCache(followers, followings int) string {
	followersStr := strconv.FormatInt(int64(followers), 10)
	followingsStr := strconv.FormatInt(int64(followings), 10)
	return fmt.Sprintf("%s,%s", followersStr, followingsStr)
}

type Relation struct {
	gorm.Model
	Uid          int64
	Fid          int64
	Relationship int8
}

type RelationCount struct {
	gorm.Model
	Uid       int64
	Follower  int
	Following int
}

type CacheRelationshipAttribute struct {
	Fid          int64     `json:"fid"`
	Relationship string    `json:"relationship"`
	UpdateTime   time.Time `json:"update_time"`
}
