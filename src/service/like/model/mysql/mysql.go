package mysql

import (
	"errors"
	"fmt"
	gomysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"nichebox/service/like/model"
	"time"
)

type MysqlInterface struct {
	db *gorm.DB
}

func (m *MysqlInterface) GetLikeByUpdateDateDesc(typeID int, uid int64, limit int, offset int) ([]*model.Like, error) {
	likeList := make([]*model.Like, 0)
	result := m.db.Model(&model.Like{}).Where("uid = ? AND type_id = ?", uid, typeID).Order("updated_at desc").Limit(limit).Offset(offset).Find(&likeList)
	if result.Error != nil {
		return nil, result.Error
	}
	return likeList, nil
}

func (m *MysqlInterface) DeleteLikeAndUpdateLikeCountTX(likeModel *model.Like) error {
	tx := m.db.Begin()
	// check if record existed or deleted
	likeModelInDB := model.Like{}
	result := tx.Unscoped().Where("uid = ? AND type_id = ? AND message_id = ?", likeModel.Uid, likeModel.TypeID, likeModel.MessageID).First(&likeModelInDB)
	if result.Error != nil {
		// not exists or other errors
		tx.Rollback()
		return result.Error
	}
	deleted := likeModelInDB.DeletedAt.Valid
	if deleted {
		// deleted, no need to delete again
		tx.Commit()
		return nil
	}
	// not deleted, now begin delete
	result = tx.Where("uid = ? AND type_id = ? AND message_id = ?", likeModel.Uid, likeModel.TypeID, likeModel.MessageID).Delete(&model.Like{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	likeCountModel := model.LikeCount{}
	result = tx.Where("message_id = ? AND type_id = ?", likeModel.MessageID, likeModel.TypeID).First(&likeCountModel)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	likeCountModel.Count -= 1
	result = tx.Model(&likeCountModel).Where("message_id = ? AND type_id = ?", likeCountModel.MessageID, likeCountModel.TypeID).Update("count", likeCountModel.Count)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

func (m *MysqlInterface) CreateLikeAndUpdateLikeCountTX(likeModel *model.Like) error {
	tx := m.db.Begin()
	// check if like record had been deleted
	needCreate := false
	likeModelInDB := model.Like{}
	result := tx.Unscoped().Where("uid = ? AND type_id = ? AND message_id = ?", likeModel.Uid, likeModel.TypeID, likeModel.MessageID).First(&likeModelInDB)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// other errors
			tx.Rollback()
			return result.Error
		}

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			needCreate = true
		}
	}

	if needCreate {
		// need to create like record
		result = tx.Create(likeModel)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	} else {
		// no need to create like record, check if this record deleted
		deleted := likeModelInDB.DeletedAt.Valid
		if deleted {
			// deleted, recover the deleted record
			result = tx.Unscoped().Model(&likeModelInDB).Update("deleted_at", nil)
			if result.Error != nil {
				tx.Rollback()
				return result.Error
			}
		} else {
			// not deleted, duplicate error
			tx.Rollback()
			mysqlErr := &gomysql.MySQLError{Number: 1062}
			return mysqlErr
		}
	}

	// check if like count record exists
	likeCountModel := model.LikeCount{}
	result = tx.Where("message_id = ? AND type_id = ?", likeModel.MessageID, likeModel.TypeID).First(&likeCountModel)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// other errors
			tx.Rollback()
			return result.Error
		}

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// like count record not found, create record now
			lk := model.LikeCount{
				MessageID: likeModel.MessageID,
				TypeID:    likeModel.TypeID,
				Count:     1,
			}
			result = tx.Create(&lk)
			if result.Error != nil {
				tx.Rollback()
				return result.Error
			}
			tx.Commit()
			return nil
		}
	}
	// like count record existed, just increase like count
	likeCountModel.Count += 1
	result = tx.Model(&likeCountModel).Where("message_id = ? AND type_id = ?", likeCountModel.MessageID, likeCountModel.TypeID).Update("count", likeCountModel.Count)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil

}

func (m *MysqlInterface) GetLikeCount(likeCountModel *model.LikeCount) error {
	result := m.db.Where("message_id = ? AND type_id = ?", likeCountModel.MessageID, likeCountModel.TypeID).First(likeCountModel)
	return result.Error
}

func NewMysqlInterface(database, username, password, host, port string, maxIdleConns, maxOpenConns, connMaxLifeTime int) (model.LikeInterface, error) {
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
	m.db.AutoMigrate(&model.Like{})
	m.db.AutoMigrate(&model.LikeCount{})
}
