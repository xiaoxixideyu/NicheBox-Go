package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"nichebox/service/post/model"
	"time"
)

type MysqlInterface struct {
	db *gorm.DB
}

func NewMysqlInterface(database, username, password, host, port string, maxIdleConns, maxOpenConns, connMaxLifeTime int) (model.PostInterface, error) {
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
	m.db.AutoMigrate(&model.Post{})
}

func (m *MysqlInterface) CreatePost(post *model.Post) error {
	result := m.db.Create(post)
	return result.Error
}

func (m *MysqlInterface) DeletePost(post *model.Post) error {
	result := m.db.Where("post_id = ?", post.PostID).Delete(post)
	return result.Error
}

func (m *MysqlInterface) GetPostByID(postID int64) (*model.Post, error) {
	var post model.Post
	result := m.db.Where("post_id = ?", postID).First(&post)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}
