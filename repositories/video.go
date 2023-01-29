package repositories

import (
	"douyin/model"
	"gorm.io/gorm"
)

var VideoRepository = newVideoRepository()

func newVideoRepository() *videoRepository {
	return &videoRepository{}
}

type videoRepository struct {
}

func (r *videoRepository) Find(db *gorm.DB, i int64) (list []model.Video) {
	db.Table("video").Where("author_id = ?", i).Order("create_time desc").Find(&list)
	return
}

func (r *videoRepository) Create(db *gorm.DB, t *model.Video) (err error) {
	err = db.Table("video").Create(t).Error
	return
}

func (r *videoRepository) FindByTime(db *gorm.DB, t string) (videos []model.Video) {
	if t != "0" {
		db.Table("video").Where("create_time >= ?", t).Order("create_time desc").Limit(30).Find(&videos)
	} else {
		db.Table("video").Order("create_time desc").Limit(30).Find(&videos)
	}
	return
}

func (r *videoRepository) IncFavoriteCount(db *gorm.DB, vid int64) (err error) {
	err = db.Table("video").Where(map[string]interface{}{
		"id": vid,
	}).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
	return
}

func (r *videoRepository) DecFavoriteCount(db *gorm.DB, vid int64) (err error) {
	err = db.Table("video").Where(map[string]interface{}{
		"id": vid,
	}).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
	return
}

func (r *videoRepository) GetVideoById(db *gorm.DB, id int64) (ret model.Video) {
	db.Table("video").Where(map[string]interface{}{
		"id": id,
	}).First(&ret)
	return
}
