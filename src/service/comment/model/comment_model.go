package model

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	CommentID       int64
	SubjectID       int64
	RootID          int64
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
