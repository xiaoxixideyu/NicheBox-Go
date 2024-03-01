package model

import "gorm.io/gorm"

type BoxUserInterface interface {
	GetTx() *gorm.DB
	IsBoxExistsByTx(*BoxUser, *gorm.DB) (bool, error)
	IsOwnerExists(*BoxUser) (bool, error)
	AddBoxUserByTx(*BoxUser, *gorm.DB) error
	RemoveBoxUserByTx(*BoxUser, *gorm.DB) error
	GetBoxUser(bid, uid int64) (*BoxUser, error)
	AddBoxUser(*BoxUser) error
	RemoveBoxUser(*BoxUser) error
	UpdateRole(*BoxUser) error
}
