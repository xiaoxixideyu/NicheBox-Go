package model

import (
	"gorm.io/gorm"
)

type BoxInterface interface {
	IsBoxExistsByTx(*Box, *gorm.DB) (bool, error)
	CreateBoxByTx(*Box, *gorm.DB) error
	RemoveBoxByTx(*Box, *gorm.DB) error
	GetTx() *gorm.DB
}
