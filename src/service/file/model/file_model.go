package model

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	FileId int64 `gorm:"index:idx_file_image_id,unique"`
	Url    string
}
