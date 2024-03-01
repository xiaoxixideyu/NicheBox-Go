package mysql

import (
	"errors"
	"fmt"
	"nichebox/service/box_user/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlInterface struct {
	db *gorm.DB
}

func NewMysqlInterface(database, username, password, host, port string, maxIdleConns, maxOpenConns, connMaxLifeTime int) (model.BoxUserInterface, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		username,
		password,
		host,
		port,
		database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to open mysql")
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("connect mysql server failed, err:" + err.Error())
		return nil, err
	}
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(connMaxLifeTime))

	m := &MysqlInterface{
		db: db,
	}
	m.autoMigrate()
	return m, nil
}

func (m *MysqlInterface) autoMigrate() {
	m.db.AutoMigrate(&model.BoxUser{})
}

func (m *MysqlInterface) GetTx() *gorm.DB {
	return m.db.Begin()
}

func (m *MysqlInterface) IsBoxExistsByTx(boxUser *model.BoxUser, tx *gorm.DB) (bool, error) {
	result := tx.Where("bid = ?", boxUser.Bid).First(&model.BoxUser{})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func (m *MysqlInterface) IsOwnerExists(boxUser *model.BoxUser) (bool, error) {
	result := m.db.Where("bid = ? and uid = ? and role = ?", boxUser.Bid, boxUser.Uid, boxUser.Role).First(&model.BoxUser{})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func (m *MysqlInterface) AddBoxUserByTx(boxUser *model.BoxUser, tx *gorm.DB) error {
	result := tx.Create(boxUser)
	return result.Error
}

func (m *MysqlInterface) RemoveBoxUserByTx(boxUser *model.BoxUser, tx *gorm.DB) error {
	result := tx.Where("bid = ? and uid = ?", boxUser.Bid, boxUser.Uid).Delete(&model.BoxUser{})
	return result.Error
}

func (m *MysqlInterface) GetBoxUser(bid, uid int64) (*model.BoxUser, error) {
	boxUser := &model.BoxUser{}
	result := m.db.Where("bid = ? and uid = ?", bid, uid).First(&boxUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return boxUser, nil
}

func (m *MysqlInterface) AddBoxUser(boxUser *model.BoxUser) error {
	tx := m.db.Begin()
	result := tx.Where("bid = ? and uid = ?", boxUser.Bid, boxUser.Uid).First(&model.BoxUser{})
	if result.Error == nil || result.RowsAffected > 0 {
		tx.Rollback()
		return model.ErrBoxUserExisted
	}
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return result.Error
	}
	result = tx.Create(boxUser)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

func (m *MysqlInterface) RemoveBoxUser(boxUser *model.BoxUser) error {
	tx := m.db.Begin()
	result := tx.Where("bid = ? and uid = ?", boxUser.Bid, boxUser.Uid).First(&model.BoxUser{})
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil
	}
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	result = tx.Where("bid = ? and uid = ?", boxUser.Bid, boxUser.Uid).Delete(&model.BoxUser{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

func (m *MysqlInterface) UpdateRole(boxUser *model.BoxUser) error {
	tx := m.db.Begin()
	result := tx.Where("bid = ? and uid = ?", boxUser.Bid, boxUser.Uid).First(&model.BoxUser{})
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil
	}
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	result = tx.Model(&model.BoxUser{}).Where("bid = ? and uid = ?", boxUser.Bid, boxUser.Uid).Update("role", boxUser.Role)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}
