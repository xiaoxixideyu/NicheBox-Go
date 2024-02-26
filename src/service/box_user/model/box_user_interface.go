package model

import "gorm.io/gorm"

type BoxUserInterface interface {
	GetTx() *gorm.DB
	IsOwnerExistsByTx(*BoxUser, *gorm.DB) (bool, error)
	AddBoxUserByTx(*BoxUser, *gorm.DB) error
	RemoveBoxUserByTx(*BoxUser, *gorm.DB) error
}
