package model

import "gorm.io/gorm"

type BoxUser struct {
	gorm.Model
	Bid  int64 `gorm:"not null"`
	Uid  int64 `gorm:"not null"`
	Role int   `gorm:"not null"`
}
