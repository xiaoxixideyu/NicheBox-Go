package model

type BoxUserCacheInterface interface {
	SetBoxUser(boxUser *BoxUserCache, expireTime int) error
	GetBoxUser(bid, uid int64, expireTime int) (*BoxUserCache, error)
	RemoveBoxUser(bid, uid int64) error
}
