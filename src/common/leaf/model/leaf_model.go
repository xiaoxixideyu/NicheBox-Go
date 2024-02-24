package model

import "gorm.io/gorm"

type Leaf struct {
	gorm.Model
	BizTag string `gorm:"not null;index:idx_leaf_biztag"`
	MaxId  int64  `gorm:"not null;default:0"`
}
