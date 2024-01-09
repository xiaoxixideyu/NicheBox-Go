package mysql

import (
	"errors"
	"fmt"
	"nichebox/common/snowflake"
	"nichebox/service/user/model"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlInterface struct {
	db *gorm.DB

	mu     sync.Mutex
	txMaps map[int64]*gorm.DB
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
		db:     db,
		txMaps: map[int64]*gorm.DB{},
	}
	m.autoMigrate()
	return m, nil
}

func (m *MysqlInterface) autoMigrate() {
	m.db.AutoMigrate(&model.User{})
}

func (m *MysqlInterface) BeginTX() (int64, error) {
	tx := m.db.Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}
	id := snowflake.GenID()

	m.mu.Lock()
	defer m.mu.Unlock()
	m.txMaps[id] = tx

	return id, nil
}

func (m *MysqlInterface) CommitTX(txId int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	tx, ok := m.txMaps[txId]
	if !ok {
		return errors.New(fmt.Sprintf("tx:%d not found", txId))
	}

	result := tx.Commit()
	if result.Error != nil {
		return result.Error
	}
	delete(m.txMaps, txId)

	return nil
}

func (m *MysqlInterface) RollbackTX(txId int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	tx, ok := m.txMaps[txId]
	if !ok {
		return errors.New(fmt.Sprintf("tx:%d not found", txId))
	}

	result := tx.Rollback()
	if result.Error != nil {
		return result.Error
	}
	delete(m.txMaps, txId)

	return nil
}

func (m *MysqlInterface) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	result := m.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (m *MysqlInterface) GetUserByUid(uid int64) (*model.User, error) {
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

func (m *MysqlInterface) UpdatePasswordByEmail(email, password string) error {
	var user model.User
	result := m.db.Model(&user).Where("email = ?", email).Update("password", password)
	return result.Error
}
