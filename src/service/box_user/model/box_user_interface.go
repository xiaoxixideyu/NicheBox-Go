package model

import "gorm.io/gorm"

type BoxUserInterface interface {
	GetTx() *gorm.DB
	IsBoxExistsByTx(*BoxUser, *gorm.DB) (bool, error)
	IsOwnerExists(*BoxUser) (bool, error)
	AddBoxUserByTx(*BoxUser, *gorm.DB) error
	RemoveBoxUserByTx(*BoxUser, *gorm.DB) error
}
