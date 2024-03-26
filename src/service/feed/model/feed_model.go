package model

import (
	"gorm.io/gorm"
	"time"
)

type Feed struct {
	gorm.Model
	FeedID      int64
	MessageID   int64
	TypeID      int
	AuthorID    int64
	PublishTime time.Time
}
