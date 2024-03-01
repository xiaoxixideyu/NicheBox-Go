package model

type BoxUserCache struct {
	Bid   int64
	Uid   int64
	Exist bool
	Role  int
}

type BoxUserCacheVal struct {
	Exist bool `json:"exist"`
	Role  int  `json:"role"`
}
