package mysql

import (
	"errors"
	"fmt"
	"nichebox/service/box_info/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlInterface struct {
	db *gorm.DB
}

func NewMysqlInterface(database, username, password, host, port string, maxIdleConns, maxOpenConns, connMaxLifeTime int) (model.BoxInterface, error) {
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
	m.db.AutoMigrate(&model.Box{})
}

func (m *MysqlInterface) IsBoxExistsByTx(box *model.Box, tx *gorm.DB) (bool, error) {
	result := tx.Where("bid = ?", box.Bid).First(&model.Box{})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func (m *MysqlInterface) CreateBoxByTx(box *model.Box, tx *gorm.DB) error {
	result := tx.Create(box)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *MysqlInterface) RemoveBoxByTx(box *model.Box, tx *gorm.DB) error {
	result := tx.Where("bid = ?", box.Bid).Delete(&model.Box{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *MysqlInterface) GetTx() *gorm.DB {
	return m.db.Begin()
}
