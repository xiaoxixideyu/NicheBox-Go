package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"nichebox/service/post/model"
	"nichebox/service/post/model/dto"
	"time"
)

type MysqlInterface struct {
	db *gorm.DB
}

func (m *MysqlInterface) GetModifiedPosts(from time.Time, to time.Time) ([]*dto.NewPostInfo, []*dto.DeletedPostInfo, error) {
	// new info
	newInfos := make([]*dto.NewPostInfo, 0, 0)
	subQuery1 := m.db.Model(&model.Post{}).Select("post_id, created_at, box_id").Where("created_at < ? AND created_at >= ?", to, from)
	subQuery2 := m.db.Model(&model.Post{}).Select("box_id, count(*) as count").Where("created_at < ? AND created_at >= ?", to, from).Group("box_id")
	result := m.db.Debug().Table("(?) as p, (?) as c", subQuery1, subQuery2).Select("p.post_id, p.created_at, c.box_id, c.count").Where("p.box_id = c.box_id").Order("box_id").Find(&newInfos)
	if result.Error != nil {
		return nil, nil, result.Error
	}

	// deleted info
	deletedInfos := make([]*dto.DeletedPostInfo, 0, 0)
	subQuery3 := m.db.Unscoped().Model(&model.Post{}).Select("post_id, deleted_at, box_id").Where("deleted_at < ? AND deleted_at >= ?", to, from)
	subQuery4 := m.db.Unscoped().Model(&model.Post{}).Select("box_id, count(*) as count").Where("deleted_at < ? AND deleted_at >= ?", to, from).Group("box_id")
	result = m.db.Debug().Unscoped().Table("(?) as p, (?) as c", subQuery3, subQuery4).Select("p.post_id, p.deleted_at, c.box_id, c.count").Where("p.box_id = c.box_id").Order("box_id").Find(&deletedInfos)
	if result.Error != nil {
		return nil, nil, result.Error
	}
	return newInfos, deletedInfos, nil
}

func NewMysqlInterface(database, username, password, host, port string, maxIdleConns, maxOpenConns, connMaxLifeTime int) (model.PostInterface, error) {
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

func (m *MysqlInterface) autoMigrate() {
	m.db.AutoMigrate(&model.Post{})
}

func (m *MysqlInterface) CreatePost(post *model.Post) error {
	result := m.db.Debug().Create(post)
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
