package dto

type DeliverFeedToOutboxMessage struct {
	FeedID      int64 `json:"feed_id"`
	AuthorID    int64 `json:"author_id"`
	MessageID   int64 `json:"message_id"`
	MessageType int   `json:"message_type"`
}
