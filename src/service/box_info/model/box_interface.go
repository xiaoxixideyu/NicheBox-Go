package model

import (
	"gorm.io/gorm"
)

type BoxInterface interface {
	GetTx() *gorm.DB
	IsBoxExistsByTx(*Box, *gorm.DB) (bool, error)
	CreateBoxByTx(*Box, *gorm.DB) error
	RemoveBoxByTx(*Box, *gorm.DB) error
	UpdateBoxByTx(*Box, *gorm.DB) error
	GetBoxInfo(*Box) error
}
