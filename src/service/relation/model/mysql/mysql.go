package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"nichebox/common/biz"
	"nichebox/service/relation/model"
	"time"
)

type MysqlInterface struct {
	db *gorm.DB
}

func (m *MysqlInterface) GetRelationship(uid int64, fid int64) (*model.Relation, error) {
	r := model.Relation{}
	rs := m.db.Model(&model.Relation{}).Where("uid = ? AND fid = ?", uid, fid).First(&r)
	if rs.Error != nil {
		return nil, rs.Error
	}
	return &r, nil
}

func (m *MysqlInterface) AddFollow(uid int64, fid int64) error {
	tx := m.db.Begin()
	rUid := model.Relation{}
	rFid := model.Relation{}
	// check if uid had relation before
	rs := tx.Model(&model.Relation{}).Where("uid = ? AND fid = ?", uid, fid).Find(&rUid)
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error
	}
	uidRelation := rs.RowsAffected != 0
	// check if fid had relation before
	rs = tx.Model(&model.Relation{}).Where("uid = ? AND fid = ?", fid, uid).Find(&rFid)
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error
	}
	fidRelation := rs.RowsAffected != 0

	// duplicated follow is not allowed
	if uidRelation && rUid.Relationship == model.RelationFriend || rUid.Relationship == model.RelationFollow {
		return nil
	}

	var newRelation int8

	if !uidRelation && !fidRelation {
		newRelation = model.RelationFollow
		// create uid
		newR := model.Relation{
			Uid:          uid,
			Fid:          fid,
			Relationship: newRelation,
		}
		rs = createRelation(tx, newR)
		if rs.Error != nil {
			tx.Rollback()
			return rs.Error
		}

	} else if !uidRelation && fidRelation {
		if rFid.Relationship == model.RelationFollow {
			newRelation = model.RelationFriend
		} else {
			newRelation = model.RelationFollow
		}

		// create uid
		newR := model.Relation{
			Uid:          uid,
			Fid:          fid,
			Relationship: newRelation,
		}
		rs = createRelation(tx, newR)
		if rs.Error != nil {
			tx.Rollback()
			return rs.Error
		}
		// update fid
		if newRelation == model.RelationFriend {
			rs = updateRelation(tx, fid, uid, newRelation)
			if rs.Error != nil {
				tx.Rollback()
				return rs.Error
			}
		}

	} else if uidRelation && !fidRelation {
		newRelation = model.RelationFollow
		// update uid
		rs = updateRelation(tx, uid, fid, newRelation)
		if rs.Error != nil {
			tx.Rollback()
			return rs.Error
		}

	} else if uidRelation && fidRelation {
		if rFid.Relationship == model.RelationFollow {
			newRelation = model.RelationFriend
		} else {
			newRelation = model.RelationFollow
		}

		// update uid
		rs = updateRelation(tx, uid, fid, newRelation)
		if rs.Error != nil {
			tx.Rollback()
			return rs.Error
		}
		// update fid
		rs = updateRelation(tx, fid, uid, newRelation)
		if rs.Error != nil {
			tx.Rollback()
			return rs.Error
		}
	}

	// update relation count
	rs = updateFollowingCount(tx, uid, 1)
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error
	}
	rs = updateFollowerCount(tx, fid, 1)
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error
	}

	tx.Commit()
	return nil
}

func (m *MysqlInterface) RemoveFollow(uid int64, fid int64) error {
	tx := m.db.Begin()
	rUid := model.Relation{}
	rFid := model.Relation{}
	// check if uid had relation before
	rs := tx.Model(&model.Relation{}).Where("uid = ? AND fid = ?", uid, fid).Find(&rUid)
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error
	}
	uidRelation := rs.RowsAffected != 0
	// uid never followed fid before
	if !uidRelation || rUid.Relationship != model.RelationFollow && rUid.Relationship != model.RelationFriend {
		tx.Commit()
		return gorm.ErrRecordNotFound
	}
	// check if fid had relation before
	rs = tx.Model(&model.Relation{}).Where("uid = ? AND fid = ?", fid, uid).Find(&rFid)
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error
	}
	fidRelation := rs.RowsAffected != 0

	// update uid
	rs = updateRelation(tx, uid, fid, model.RelationNone)

	// update fid
	if fidRelation && rFid.Relationship == model.RelationFriend {
		rs = updateRelation(tx, fid, uid, model.RelationFollow)
		if rs.Error != nil {
			tx.Rollback()
			return rs.Error
		}
	}

	// update relation count
	rs = updateFollowingCount(tx, uid, -1)
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error
	}
	rs = updateFollowerCount(tx, fid, -1)
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error
	}

	tx.Commit()
	return nil
}

func (m *MysqlInterface) GetFollowers(uid int64, page, size int, order string) ([]*model.Relation, error) {
	relations := make([]*model.Relation, 0, 10)

	var rs *gorm.DB
	var orderExpr string
	switch order {
	case biz.OrderByCreateTimeAsc:
		orderExpr = "updated_at asc"
	case biz.OrderByCreateTimeDesc:
		orderExpr = "updated_at desc"
	default:
		orderExpr = "updated_at desc"
	}

	if size == -1 {
		rs = m.db.Model(&model.Relation{}).Where("fid = ? AND relationship = ? OR relationship = ?", uid, model.RelationFollow, model.RelationFriend).Order(orderExpr).Find(&relations)
	} else {
		offset := (page - 1) * size
		rs = m.db.Model(&model.Relation{}).Where("fid = ? AND relationship = ? OR relationship = ?", uid, model.RelationFollow, model.RelationFriend).Order(orderExpr).Offset(offset).Limit(size).Find(&relations)
	}

	if rs.Error != nil {
		return nil, rs.Error
	}
	return relations, nil
}

func (m *MysqlInterface) GetFollowings(uid int64, page, size int, order string) ([]*model.Relation, error) {
	relations := make([]*model.Relation, 0, 10)

	var rs *gorm.DB
	var orderExpr string
	switch order {
	case biz.OrderByCreateTimeAsc:
		orderExpr = "updated_at asc"
	case biz.OrderByCreateTimeDesc:
		orderExpr = "updated_at desc"
	default:
		orderExpr = "updated_at desc"
	}

	if size == -1 {
		rs = m.db.Model(&model.Relation{}).Where("uid = ? AND relationship = ? OR relationship = ?", uid, model.RelationFollow, model.RelationFriend).Order(orderExpr).Find(&relations)
	} else {
		offset := (page - 1) * size
		rs = m.db.Model(&model.Relation{}).Where("uid = ? AND relationship = ? OR relationship = ?", uid, model.RelationFollow, model.RelationFriend).Order(orderExpr).Offset(offset).Limit(size).Find(&relations)
	}

	if rs.Error != nil {
		return nil, rs.Error
	}
	return relations, nil
}

func (m *MysqlInterface) GetFollowerCount(uid int64) (int, error) {
	mo := model.RelationCount{
		Uid:       uid,
		Follower:  0,
		Following: 0,
	}
	rs := firstOrCreateRelationCount(m.db, &mo)
	if rs.Error != nil {
		return 0, rs.Error
	}
	return mo.Follower, nil
}

func (m *MysqlInterface) GetFollowingCount(uid int64) (int, error) {
	mo := model.RelationCount{
		Uid:       uid,
		Follower:  0,
		Following: 0,
	}
	rs := firstOrCreateRelationCount(m.db, &mo)
	if rs.Error != nil {
		return 0, rs.Error
	}
	return mo.Following, nil
}

func NewMysqlInterface(database, username, password, host, port string, maxIdleConns, maxOpenConns, connMaxLifeTime int) (model.RelationInterface, error) {
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
	m.db.AutoMigrate(&model.Relation{})
	m.db.AutoMigrate(&model.RelationCount{})
}

func createRelation(db *gorm.DB, m model.Relation) *gorm.DB {
	return db.Model(&model.Relation{}).Create(&m)
}

func updateRelation(db *gorm.DB, uid, fid int64, newRelation int8) *gorm.DB {
	return db.Model(&model.Relation{}).Where("uid = ? AND fid = ?", uid, fid).Update("relationship", newRelation)
}

func firstOrCreateRelationCount(db *gorm.DB, m *model.RelationCount) *gorm.DB {
	return db.Model(&model.RelationCount{}).Where("uid = ?", m.Uid).FirstOrCreate(m)
}

func updateFollowerCount(db *gorm.DB, uid int64, delta int) *gorm.DB {
	rs := db.Model(&model.RelationCount{}).Where("uid = ?", uid).UpdateColumn("follower", gorm.Expr("follower + ?", delta))
	if rs.RowsAffected == 0 {
		m := model.RelationCount{Uid: uid, Follower: delta}
		return firstOrCreateRelationCount(db, &m)
	}
	return rs
}

func updateFollowingCount(db *gorm.DB, uid int64, delta int) *gorm.DB {
	rs := db.Model(&model.RelationCount{}).Where("uid = ?", uid).UpdateColumn("following", gorm.Expr("following + ?", delta))
	if rs.RowsAffected == 0 {
		m := model.RelationCount{Uid: uid, Following: delta}
		return firstOrCreateRelationCount(db, &m)
	}
	return rs
}
