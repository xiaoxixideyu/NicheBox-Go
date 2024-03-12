package model

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID              uint      `gorm:"primarykey"`
	CreatedAt       time.Time `gorm:"index:idx_root_id_comment_id_deleted_at_created_at,priority:4"`
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index;index:idx_root_id_comment_id_deleted_at_created_at,priority:3"`
	CommentID       int64          `gorm:"index:idx_comment_id;index:idx_root_id_comment_id_deleted_at_created_at,priority:2"`
	SubjectID       int64
	RootID          int64 `gorm:"index:idx_root_id_comment_id_deleted_at_created_at,priority:1"`
	ParentID        int64
	DialogID        int64
	OwnerID         int64
	LikeCount       int `gorm:"default:0"`
	Floor           int
	InnerFloorCount int
	Status          int
}

type Subject struct {
	gorm.Model
	TypeID       int   `gorm:"index:idx_like_message_type,priority:2"`
	MessageID    int64 `gorm:"index:idx_like_message_type,priority:1"`
	CommentCount int
}

type CommentContent struct {
	CommentID int64 `gorm:"primaryKey"`
	Content   string
}

type CommentCache struct {
	CommentID       int64
	SubjectID       int64
	RootID          int64
	ParentID        int64
	DialogID        int64
	OwnerID         int64
	LikeCount       int
	Floor           int
	InnerFloorCount int
	Status          int
	Content         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
