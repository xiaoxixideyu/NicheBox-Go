package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	PostID   int64 `gorm:"index:idx_post_pid"`
	AuthorID int64
	BoxID    int64
	Title    string
	Content  string
	Photos   string `gorm:"size:98"`
	Cover    string `gorm:"size:10"`
	UserView int    `gorm:"default:0"`
}
