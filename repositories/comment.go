package repositories

import (
	"douyin/model"
	"douyin/util"
	"gorm.io/gorm"
	"time"
)

var CommentRepository = newCommentRepository()

func newCommentRepository() *commentRepository {
	return &commentRepository{}
}

type commentRepository struct {
}

func (r commentRepository) Publish(db *gorm.DB, uid int64, vid int64, t string) (err error) {
	c := model.Comment{
		Model:      model.Model{
			Id: util.GetOnlyId(),
		},
		VideoId:    vid,
		UserId:     uid,
		Content:    t,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		State:      true,
	}
	err = db.Table("comment").Create(&c).Error
	return
}

func (r commentRepository) Delete(db *gorm.DB, cid int64) error {
	return db.Table("comment").Where("id = ?", cid).Update("state", false).Error
}

func (r commentRepository) FindByTimeDesc(db *gorm.DB, vid int64) (ret []model.Comment) {
	db.Table("comment").Where(&model.Comment{
		VideoId: vid,
		State:   true,
	}).Order("create_time desc").Find(&ret)
	return ret
}

func (r commentRepository) IncComment(db *gorm.DB, vid int64) (err error) {
	err = db.Table("video").Where(map[string]interface{}{
		"id": vid,
	}).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error
	return
}

func (r commentRepository) DecComment(db *gorm.DB, vid int64) (err error) {
	err = db.Table("video").Where(map[string]interface{}{
		"id": vid,
	}).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error
	return
}

func (r commentRepository) GetCommentById(db *gorm.DB, id int64) (ret model.Comment) {
	db.Table("comment").Where(map[string]interface{}{
		"id": id,
	}).First(&ret)
	return
}
