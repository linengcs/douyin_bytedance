package services

import (
	"douyin/model"
	"douyin/repositories"
	"douyin/util"
	"errors"
	"gorm.io/gorm"
	"log"
	"strconv"
	"strings"
)

var UserService = newUserService()

func newUserService() *userService {
	return &userService{}
}

type userService struct {
}

func (s *userService) GetUserInfoById(cId int64, id int64) *model.UserInfo {
	userInfo := repositories.UserRepository.GetUserInfoById(Db, id)

	result, _ := RedisConn.Get("user_follow" + strconv.Itoa(int(cId)) + ":" + strconv.Itoa(int(id))).Result()
	if result == "1" {
		userInfo.IsFollow = true
	}
	return userInfo
}

func (s *userService) GetByName(username string) *model.User {
	return repositories.UserRepository.GetByUsername(Db, username)
}

func (s *userService) SignUp(username, password string) *model.User {
	username = strings.TrimSpace(username)
	user := &model.User{
		Model: model.Model{
			Id: util.GetOnlyId(),
		},
		Name:          username,
		Password:      EncryptService.hashAndSalt([]byte(password)),
	}

	err := repositories.UserRepository.Create(Db, user)
	if err != nil {
		log.Println(err)
		return nil
	}
	return user
}

func (s *userService) SignIn(username, password string) (*model.User, error) {
	user := repositories.UserRepository.GetByUsername(Db, username)
	ok := EncryptService.comparePasswords(user.Password, []byte(password))
	if ok == false {
		return nil, errors.New("check password")
	}
	return user, nil
}

func (s *userService) GetFollows(userId int64) (info []model.UserFollow) {
	info = repositories.UserRepository.GetFollows(Db, userId)
	return
}

func (s *userService) GetFollowers(userId int64) (info []model.UserFollow) {
	info = repositories.UserRepository.GetFollowers(Db, userId)
	return
}

func (s *userService) FollowUser(uid int64, tid int64, actionType int) (status bool, err error) {
	switch actionType {
	case 1:
		err = Db.Transaction(func(tx *gorm.DB) error {
			if err := repositories.UserRepository.FollowUser(Db, uid, tid); err != nil {
				log.Println(err)
				return err
			}
			if err = repositories.UserRepository.IncrFollow(Db, uid); err != nil {
				log.Println(err)
				return err
			}
			if err = repositories.UserRepository.IncrFollower(Db, tid); err != nil {
				log.Println(err)
				return err
			}
			RedisConn.Set("user_follow" + strconv.Itoa(int(uid)) + ":" + strconv.Itoa(int(tid)), 1, 0)
			return nil
		})
	case 2:
		err = Db.Transaction(func(tx *gorm.DB) error {
			if err = repositories.UserRepository.UnFollowUser(Db, uid, tid); err != nil {
				log.Println(err)
				return err
			}
			if err = repositories.UserRepository.DecrFollow(Db, uid); err != nil {
				log.Println(err)
				return err
			}
			if err = repositories.UserRepository.DecrFollower(Db, tid); err != nil {
				log.Println(err)
				return err
			}
			RedisConn.Set("user_follow"+strconv.Itoa(int(uid)) + ":" + strconv.Itoa(int(tid)), 0, 0)
			return nil
		})
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *userService) LikeVideo(uid int64, vid int64, actionType int) (err error) {
	switch actionType {
	case 1:
		err = Db.Transaction(func(tx *gorm.DB) error {
			if err = repositories.UserRepository.LikeVideo(Db, uid, vid); err != nil {
				log.Println(err)
				return err
			}
			if err = repositories.VideoRepository.IncFavoriteCount(Db, vid); err != nil {
				log.Println(err)
				return err
			}
			return nil
		})

	case 2:
		err = Db.Transaction(func(tx *gorm.DB) error {
			if err = repositories.UserRepository.UnLikeVideo(Db, uid, vid); err != nil {
				log.Println(err)
				return err
			}
			if err = repositories.VideoRepository.DecFavoriteCount(Db, vid); err != nil {
				log.Println(err)
				return err
			}
			return nil
		})
	}
	return err
}

func (s *userService) CheckUsername(n string) bool {
	return repositories.UserRepository.CheckUsername(Db, n)
}

func (s *userService) IsFavorite(id int64, vid int64) bool {
	return repositories.UserRepository.IsFavorite(Db, id, vid)
}

func (s *userService) GetUserById(id int64) *model.User{
	return repositories.UserRepository.Get(Db, id)
}