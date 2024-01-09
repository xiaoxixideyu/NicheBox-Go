package mysql

import (
	"fmt"
	"nichebox/service/user/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlInterface struct {
	db *gorm.DB
}

func NewMysqlInterface(database, username, password, host, port string, maxIdleConns, maxOpenConns, connMaxLifeTime int) (model.UserInterface, error) {
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
	m.db.AutoMigrate(&model.User{})
}

func (m *MysqlInterface) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	result := m.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (m *MysqlInterface) GerUserByUid(uid int64) (*model.User, error) {
	var user model.User
	result := m.db.Where("uid = ?", uid).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (m *MysqlInterface) CreateUser(user *model.User) error {
	result := m.db.Create(user)
	return result.Error
}

func (m *MysqlInterface) UpdateUserTX(user *model.User) error {
	tx := m.db.Begin()
	result := tx.Where("uid = ?", user.Uid).Find(&model.User{})
	if result.Error != nil {
		tx.Commit()
		return result.Error
	}
	tx.Save(user)
	tx.Commit()
	return nil
}
