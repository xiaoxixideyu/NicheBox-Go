package dto

import "time"

type RebuildCacheSubjectCommentIndexMessage struct {
	CreateDate time.Time `json:"create_date"`
	SubjectID  int64     `json:"subject_id"`
}

type RebuildCacheInnerFloorCommentIndexMessage struct {
	RootIDs []int64 `json:"root_ids"`
}

type UpdateCommentLikeCountMessage struct {
	CommentID int64 `json:"comment_id"`
	Delta     int   `json:"delta"`
}

type CommentNotificationMessage struct {
	NewCommentOwner int64  `json:"new_comment_owner"`
	Info            string `json:"info"`
}
