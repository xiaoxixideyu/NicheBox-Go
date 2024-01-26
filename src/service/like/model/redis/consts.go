package redis

const (
	Separator           = ":"
	HistoryListLength   = 10
	LikeCountExpiration = 3 * 24 * 60 * 60

	KeyPrefixLike = "like:"

	KeyCount   = "count:"
	KeyHistory = "history:"
)
