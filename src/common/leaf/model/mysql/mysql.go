package mysql

import (
	"errors"
	"fmt"
	"nichebox/common/leaf/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlInterface struct {
	db *gorm.DB
}

func NewMysqlInterface(database, username, password, host, port string, maxIdleConns, maxOpenConns, connMaxLifeTime int) (model.LeafInterface, error) {
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
	m.db.AutoMigrate(&model.Leaf{})
}

func (m *MysqlInterface) CreateLeafTX(leaf *model.Leaf) error {
	var temp model.Leaf
	tx := m.db.Begin()
	result := tx.Where("biz_tag = ?", leaf.BizTag).First(&temp)
	if result.Error == nil && result.RowsAffected > 0 {
		tx.Rollback()
		return errors.New("leaf exist")
	}
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return result.Error
	}
	if result := tx.Create(leaf); result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

func (m *MysqlInterface) NextTX(bizTag string, step int) (int64, int64, error) {
	var leaf model.Leaf
	tx := m.db.Begin()
	result := tx.Where("biz_tag = ?", bizTag).First(&leaf)
	if result.Error != nil {
		tx.Rollback()
		return 0, 0, result.Error
	}
	start := leaf.MaxId + 1
	leaf.MaxId += int64(step)
	end := leaf.MaxId
	result = tx.Save(&leaf)
	if result.Error != nil {
		tx.Rollback()
		return 0, 0, result.Error
	}
	tx.Commit()
	return start, end, nil
}
