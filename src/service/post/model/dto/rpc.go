package dto

import "time"

type NewPostInfo struct {
	BoxID     int64     `gorm:"box_id"`
	Count     int       `gorm:"count"`
	PostID    int64     `gorm:"post_id"`
	CreatedAt time.Time `gorm:"created_at"`
}

type DeletedPostInfo struct {
	BoxID     int64     `gorm:"box_id"`
	Count     int       `gorm:"count"`
	PostID    int64     `gorm:"post_id"`
	DeletedAt time.Time `gorm:"deleted_at"`
}
