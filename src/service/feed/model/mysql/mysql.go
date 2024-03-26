package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"nichebox/service/feed/model"
	"time"
)

type MysqlInterface struct {
	db *gorm.DB
}

func (m *MysqlInterface) GetFeeds(followings []int64, page, size int) ([]*model.Feed, error) {
	feeds := make([]*model.Feed, 0, size)
	offset := (page - 1) * size
	rs := m.db.Model(&model.Feed{}).Where("author_id IN ?", followings).Order("publish_time desc").Offset(offset).Limit(size).Find(&feeds)
	if rs.Error != nil {
		return nil, rs.Error
	}
	return feeds, nil

}

func (m *MysqlInterface) AddFeed(feed *model.Feed) error {
	rs := m.db.Model(&model.Feed{}).Create(feed)
	if rs.Error != nil {
		return rs.Error
	}
	return nil
}

func (m *MysqlInterface) autoMigrate() {
	m.db.AutoMigrate(&model.Feed{})
}

func NewMysqlInterface(database, username, password, host, port string, maxIdleConns, maxOpenConns, connMaxLifeTime int) (model.FeedInterface, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
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
