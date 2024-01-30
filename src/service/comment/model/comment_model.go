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
	Floor           int `gorm:"autoIncrement"`
	InnerFloorCount int
	Status          uint8
}

type Subject struct {
	gorm.Model
	TypeID       uint8 `gorm:"index:idx_like_message_type,priority:2"`
	MessageID    int64 `gorm:"index:idx_like_message_type,priority:1"`
	CommentCount int
}

type CommentContent struct {
	CommentID int64 `gorm:"primaryKey"`
	Content   string
}

type CommentInfoCache struct {
	CommentID       int64
	SubjectID       int64
	RootID          int64
	ParentID        int64
	DialogID        int64
	OwnerID         int64
	Floor           int
	InnerFloorCount int
	Status          uint8
	Content         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
