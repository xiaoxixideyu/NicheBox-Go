package model

import "gorm.io/gorm"

type Box struct {
	gorm.Model
	Bid          int64  `gorm:"not null;unique;index:idx_box_bid"`
	Name         string `gorm:"not null"`
	Introduction string
}
