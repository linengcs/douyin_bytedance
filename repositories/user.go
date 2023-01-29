package repositories

import (
	"douyin/model"
	"gorm.io/gorm"
)

var UserRepository = newUserRepository()

func newUserRepository() *userRepository {
	return &userRepository{}
}

type userRepository struct {
}

func (r *userRepository) Get(db *gorm.DB, id int64) *model.User {
	ret := &model.User{}
	db.Table("user").First(ret, "id = ?", id)
	return ret
}

func (r *userRepository) Create(db *gorm.DB, t *model.User) (err error) {
	err = db.Table("user").Create(t).Error
	return
}

func (r *userRepository) GetByUsername(db *gorm.DB, username string) *model.User {
	ret := &model.User{}
	db.Table("user").Where("name = ?", username).First(ret)
	return ret
}

func (r *userRepository) GetUserInfoById(db *gorm.DB, id int64) *model.UserInfo {
	ret := &model.UserInfo{}
	db.Table("user").Select([]string{
		"id",
		"name",
		"follow_count",
		"follower_count",
	}).Where(&model.UserInfo{
		Id: id,
	}).First(ret)
	return ret
}

func (r *userRepository) GetFollows(db *gorm.DB, uid int64) (ret []model.UserFollow) {
	db.Table("follow").Where(map[string]interface{}{
		"user_id": uid,
		"state":  true,
	}).Find(&ret)
	return ret
}

func (r *userRepository) GetFollowers(db *gorm.DB, uid int64) (ret []model.UserFollow) {
	db.Table("follow").Where(map[string]interface{}{
		"follow_id": uid,
		"state":    true,
	}).Find(&ret)
	return ret
}

func (r *userRepository) FollowUser(db *gorm.DB, uid int64, tid int64) (err error) {
	var follow model.UserFollow
	if err = db.Table("follow").Where(map[string]interface{}{
		"user_id":   uid,
		"follow_id": tid,
	}).First(&follow).Error; err != nil {
		db.Table("follow").Create(&model.UserFollow{
			UserId:   uid,
			FollowId: tid,
			State:    true,
		})
	} else {
		db.Table("follow").Where(map[string]interface{}{
			"user_id":   uid,
			"follow_id": tid,
		}).Update("state", true)
	}
	return
}

func (r *userRepository) UnFollowUser(db *gorm.DB, uid int64, tid int64) (err error) {
	err = db.Table("follow").Where(map[string]interface{}{
		"user_id":   uid,
		"follow_id": tid,
	}).Update("state", false).Error
	return
}

func (r *userRepository) GetUsers(db *gorm.DB, ids []int64) (ret []model.User) {
	db.Table("user").Where("id in (?)", ids).Find(&ret)
	return
}

func (r *userRepository) LikeVideo(db *gorm.DB, uid int64, vid int64) (err error) {
	var uv model.UserVideo
	if err = db.Table("user_video").Where(map[string]interface{}{
		"user_id":  uid,
		"video_id": vid,
	}).First(&uv).Error; err != nil {
		db.Table("user_video").Create(&model.UserVideo{
			UserId:  uid,
			VideoId: vid,
			State:   true,
		})
	} else {
		db.Table("user_video").Where(map[string]interface{}{
			"user_id":  uid,
			"video_id": vid,
		}).Update("state", true)
	}
	return
}

func (r *userRepository) UnLikeVideo(db *gorm.DB, uid int64, vid int64) (err error) {
	err = db.Table("user_video").Where(map[string]interface{}{
		"user_id":  uid,
		"video_id": vid,
	}).Update("state", false).Error
	return
}

func (r *userRepository) GetLikes(db *gorm.DB, id int64) (ret []int64) {
	db.Table("user_video").Select("video_id").Where(map[string]interface{}{
		"user_id": id,
		"state":  true,
	}).Find(&ret)
	return
}

func (r *userRepository) GetVideos(db *gorm.DB, vid []int64) (ret []model.Video) {
	db.Table("video").Where("id in (?)", vid).Find(&ret)
	return
}

func (r *userRepository) CheckUsername(db *gorm.DB, n string) bool {
	ret := &model.User{}
	db.Table("user").Where("name = ?", n).First(ret)
	return ret.Password != ""
}

func (r *userRepository) CheckPassword(db *gorm.DB, username string, password string) (*model.User, error) {
	ret := &model.User{}
	err := db.Table("user").Where(map[string]interface{}{
		"name":     username,
		"password": password,
	}).First(ret).Error
	return ret, err
}

func (r *userRepository) DecrFollow(db *gorm.DB, id int64) (err error) {
	err = db.Table("user").Where(map[string]interface{}{
		"id": id,
	}).UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1)).Error
	return
}

func (r *userRepository) IncrFollow(db *gorm.DB, id int64) (err error) {
	err = db.Table("user").Where(map[string]interface{}{
		"id": id,
	}).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error
	return
}

func (r *userRepository) IncrFollower(db *gorm.DB, id int64) (err error) {
	err = db.Table("user").Where(map[string]interface{}{
		"id": id,
	}).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error
	return
}

func (r *userRepository) DecrFollower(db *gorm.DB, id int64) (err error) {
	err = db.Table("user").Where(map[string]interface{}{
		"id": id,
	}).UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1)).Error
	return
}

func (r *userRepository) IsFavorite(db *gorm.DB, id int64, vid int64) bool {
	uv := model.UserVideo{}
	db.Table("user_video").Where(map[string]interface{}{
		"user_id":  id,
		"video_id": vid,
		"state":   true,
	}).First(&uv)
	return uv.State
}

func (r *userRepository) GetPassword(db *gorm.DB, username string) (ret string, err error) {
	err = db.Table("user").Select("password").Where(map[string]interface{}{
		"name":     username,
	}).First(&ret).Error
	return
}

func (r *userRepository) IsFollows(db *gorm.DB, uid int64, id int64) (ret model.UserFollow) {
	db.Table("follow").Where(map[string]interface{}{
		"follow_id": id,
		"user_id":   uid,
	}).Find(&ret)
	return ret
}
