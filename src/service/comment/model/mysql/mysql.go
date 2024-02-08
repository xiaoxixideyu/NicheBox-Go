package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"nichebox/common/biz"
	"nichebox/service/comment/model"
	"time"
)

type MysqlInterface struct {
	db *gorm.DB
}

func (m *MysqlInterface) GetSubjectBySubjectID(subjectID int64) (*model.Subject, error) {
	subject := model.Subject{}
	result := m.db.Model(&model.Subject{}).Where("id = ?", subjectID).First(&subject)
	if result.Error != nil {
		return nil, result.Error
	}
	return &subject, nil
}

func (m *MysqlInterface) BatchGetAllInnerFloorCommentsAndInnerFloorCounts(rootIDs []int64) ([]*model.Comment, []int, error) {
	tx := m.db.Begin()

	counts := make([]int, 0, len(rootIDs))
	result := tx.Model(&model.Comment{}).Select("inner_floor_count").Where("comment_id in ?", rootIDs).Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(comment_id,?)", Vars: []interface{}{rootIDs}, WithoutParentheses: true},
	}).Find(&counts)
	fmt.Printf("rootids:%v\ncounts:%v\n", rootIDs, counts)
	if result.Error != nil {
		tx.Rollback()
		return nil, nil, result.Error
	}

	subComments := make([]*model.Comment, 0, len(rootIDs))
	subQuery := tx.Model(&model.Comment{}).Select("*").Where("c.root_id = root_id").Order("like_count asc")
	result = tx.Table("(?) as c", m.db.Model(&model.Comment{})).Where("root_id in ? AND comment_id in ( select sub.comment_id from (?) as sub)", rootIDs, subQuery).Order("root_id asc").Find(&subComments)
	if result.Error != nil {
		tx.Rollback()
		return nil, nil, result.Error
	}

	tx.Commit()
	return subComments, counts, nil
}

func (m *MysqlInterface) GetRootCommentsBySubjectID(subjectID int64, page, size int, order string) ([]*model.Comment, error) {
	comments := make([]*model.Comment, 0)
	var result *gorm.DB
	var orderExpr string
	switch order {
	case biz.OrderByTimeAsc:
		orderExpr = "floor asc"
	case biz.OrderByTimeDesc:
		orderExpr = "floor desc"
	case biz.OrderByLikeCount:
		orderExpr = "like_count desc"
	default:
		orderExpr = "floor asc"
	}
	if size == -1 {
		result = m.db.Model(&model.Comment{}).Where("subject_id = ? AND root_id = 0", subjectID).Order(orderExpr).Find(&comments)
	} else {
		offset := (page - 1) * size
		result = m.db.Model(&model.Comment{}).Where("subject_id = ? AND root_id = 0", subjectID).Order(orderExpr).Offset(offset).Limit(size).Find(&comments)
	}
	return comments, result.Error
}

func (m *MysqlInterface) BatchGetInnerFloorComments(rootIDs []int64, page, size int) ([]*model.Comment, error) {
	offset := (page - 1) * size
	subComments := make([]*model.Comment, 0, len(rootIDs))
	if size == -1 {
		offset = 0
		size = math.MaxInt16
	}
	subQuery := m.db.Model(&model.Comment{}).Select("*").Where("c.root_id = root_id").Order("like_count asc").Offset(offset).Limit(size)
	rs := m.db.Table("(?) as c", m.db.Model(&model.Comment{})).Where("root_id in ? AND comment_id in ( select sub.comment_id from (?) as sub)", rootIDs, subQuery).Order("root_id asc").Find(&subComments)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return subComments, nil
}

func (m *MysqlInterface) GetInnerFloorCommentsAndContentsByRootID(rootID int64, page, size int) ([]*model.Comment, []*model.CommentContent, error) {
	comments := make([]*model.Comment, 0, size)
	contents := make([]*model.CommentContent, 0, size)
	offset := (page - 1) * size
	result := m.db.Model(&model.Comment{}).Where("root_id = ?", rootID).Order("updated_at asc").Offset(offset).Limit(size).Find(&comments)
	if result.Error != nil {
		return nil, nil, result.Error
	}
	ids := make([]int64, 0, len(comments))
	for _, c := range comments {
		ids = append(ids, c.CommentID)
	}
	result = m.db.Model(&model.CommentContent{}).Where("comment_id in ?", ids).Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(comment_id,?)", Vars: []interface{}{ids}, WithoutParentheses: true},
	}).Find(&contents)
	if result.Error != nil {
		return nil, nil, result.Error
	}
	return comments, contents, nil
}

func (m *MysqlInterface) BatchGetComments(ids []int64) ([]*model.Comment, error) {
	comments := make([]*model.Comment, 0, len(ids))
	result := m.db.Model(&model.Comment{}).Where("comment_id in ?", ids).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

func (m *MysqlInterface) BatchGetCommentsContents(ids []int64) ([]*model.CommentContent, error) {
	contents := make([]*model.CommentContent, 0, len(ids))
	result := m.db.Model(&model.CommentContent{}).Where("comment_id in ?", ids).Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(comment_id,?)", Vars: []interface{}{ids}, WithoutParentheses: true},
	}).Find(&contents)
	if result.Error != nil {
		return nil, result.Error
	}
	return contents, nil
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
	// clause.returning not working on mysql
	result := tx.Clauses(clause.Returning{}).Where("comment_id = ?", commentID).Delete(&comment)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	// mysql did not support return the record which just has been deleted, so we have to query
	result = tx.Unscoped().Model(&model.Comment{}).Where("comment_id").First(&comment)
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
	result := m.db.Where("message_id = ? AND type_id = ?", subject.MessageID, subject.TypeID).FirstOrCreate(subject)
	return result.Error
}

func (m *MysqlInterface) AddCommentAndUpdateSubjectTX(subject *model.Subject, comment *model.Comment, content *model.CommentContent) error {
	tx := m.db.Begin()

	if comment.RootID != 0 {
		// if this is a sub comment, its floor = 0
		comment.Floor = 0
	} else {
		// if this is a root comment, assign a floor
		var maxFloor int
		if subject.CommentCount == 0 {
			maxFloor = 1
		} else {
			result := tx.Table("comments").Select("max(floor) AS max_floor").Where("subject_id = ?", subject.ID).Find(&maxFloor)
			if result.Error != nil {
				tx.Rollback()
				return result.Error
			}
		}
		comment.Floor = maxFloor + 1
	}

	// insert comment
	result := tx.Create(comment)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	// insert comment content
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
