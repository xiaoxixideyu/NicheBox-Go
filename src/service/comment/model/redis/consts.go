package redis

const (
	Separator        = ":"
	NormalExpiration = 7 * 24 * 60 * 60
	LongExpiration   = 30 * 24 * 60 * 60

	KeyPrefixComment = "comment:"

	KeyFloor            = "floor:"
	KeyLikeCount        = "likecount:"
	KeyCreateTime       = "createtime:"
	KeyInnerFloor       = "innerfloor:"
	KeySubject          = "subject:"
	KeySubjectByMessage = "subjectbymessage:"
	KeyCommentIndex     = "commentindex:"
	KeyContent          = "content:"
)
