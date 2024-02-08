package dto

type UpdateCommentLikeCountMessage struct {
	CommentID int64 `json:"comment_id"`
	Delta     int   `json:"delta"`
}
