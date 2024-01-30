package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"nichebox/service/comment/model"
	"time"
)

type MysqlInterface struct {
	db *gorm.DB
}

func (m *MysqlInterface) GetCommentContent(commentID int64) (*model.CommentContent, error) {
	content := model.CommentContent{}
	result := m.db.Where("comment_id", commentID).First(&content)
	if result.Error != nil {
		return nil, result.Error
	}
	return &content, nil
}

func (m *MysqlInterface) GetComment(commentID int64) (*model.Comment, error) {
	comment := model.Comment{}
	result := m.db.Where("comment_id = ?", commentID).First(&comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

func (m *MysqlInterface) DeleteCommentAndUpdateSubjectTX(commentID int64) (*model.Comment, error) {
	tx := m.db.Begin()
	// delete comment
	comment := model.Comment{}
	result := tx.Clauses(clause.Returning{}).Where("comment_id = ?", commentID).Delete(&comment)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	// update subject comment count
	result = tx.Model(&model.Subject{}).Where("id = ?", comment.SubjectID).Update("comment_count", gorm.Expr("comment_count - ?", 1))
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	// update root comment inner floor count
	if comment.RootID != 0 {
		result = tx.Model(&model.Comment{}).Where("comment_id = ?", comment.RootID).Update("inner_floor_count", gorm.Expr("inner_floor_count - ?", 1))
		if result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
	}
	tx.Commit()
	return &comment, nil
}

func (m *MysqlInterface) FirstOrCreateSubject(subject *model.Subject) error {
	result := m.db.Where("message_id = ? AND message_type = ?", subject.MessageID, subject.TypeID).FirstOrCreate(subject)
	return result.Error
}

func (m *MysqlInterface) AddCommentAndUpdateSubjectTX(subject *model.Subject, comment *model.Comment, content *model.CommentContent) error {
	tx := m.db.Begin()
	// insert comment
	result := tx.Create(comment)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	// insert comment content
	content.CommentID = comment.CommentID
	result = tx.Create(content)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	// update subject comment count
	result = tx.Model(&model.Subject{}).Where("id = ?", subject.ID).Update("comment_count", gorm.Expr("comment_count + ?", 1))
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	// update root comment inner floor count
	if comment.RootID != 0 {
		result = tx.Model(&model.Comment{}).Where("comment_id = ?", comment.RootID).Update("inner_floor_count", gorm.Expr("inner_floor_count + ?", 1))
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}
	tx.Commit()
	return nil
}

func NewMysqlInterface(database, username, password, host, port string, maxIdleConns, maxOpenConns, connMaxLifeTime int) (model.CommentInterface, error) {
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
	m.db.AutoMigrate(&model.Comment{})
	m.db.AutoMigrate(&model.CommentContent{})
	m.db.AutoMigrate(&model.Subject{})
}
