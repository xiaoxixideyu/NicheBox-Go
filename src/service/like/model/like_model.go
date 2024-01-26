package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	MessageID int64 `gorm:"index:idx_like_message_uid_type,priority:3,unique"`
	Uid       int64 `gorm:"index:idx_like_message_uid_type,priority:1,unique"`
	TypeID    uint8 `gorm:"index:idx_like_message_uid_type,priority:2,unique"`
}

type LikeCount struct {
	TypeID    uint8 `gorm:"index:idx_like_message_type,priority:2"`
	MessageID int64 `gorm:"index:idx_like_message_type,priority:1"`
	Count     int
}
