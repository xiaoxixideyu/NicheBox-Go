package mysql

import (
	"fmt"
	"nichebox/service/file/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlInterface struct {
	db *gorm.DB
}

func NewMysqlInterface(database, username, password, host, port string, maxIdleConns, maxOpenConns, connMaxLifeTime int) (model.ImageFileInterface, error) {
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
	m.db.AutoMigrate(&model.Image{})
}

func (m *MysqlInterface) CreateImage(image *model.Image) error {
	result := m.db.Create(image)
	return result.Error
}

func (m *MysqlInterface) GetImageByFileId(fileId int64) (*model.Image, error) {
	var image model.Image
	result := m.db.Where("file_id = ?", fileId).Find(&image)
	if result.Error != nil {
		return nil, result.Error
	}
	return &image, nil
}
