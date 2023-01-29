package render

import (
	"douyin/model"
	"douyin/repositories"
	"douyin/services"
	"strconv"
)

func BuildUserDefaultIfNull(id int64, cId int64) *model.UserInfo {
	u := services.UserService.GetUserInfoById(cId, id)

	result, _ := services.RedisConn.Get("user_follow" + strconv.Itoa(int(cId)) + ":" + strconv.Itoa(int(id))).Result()
	if result == "1" {
		u.IsFollow = true
	}
	return u
}

func BuildUserInfo(s *model.UserFollow, t string, uid int64) *model.UserInfo {
	if s == nil {
		return nil
	}

	info := &model.UserInfo{}
	var u *model.User
	if t == "follow" {
		u = repositories.UserRepository.Get(services.Db, s.FollowId)
		ret := repositories.UserRepository.IsFollows(services.Db, uid, s.FollowId)
		info.IsFollow = ret.State
	} else {
		u = repositories.UserRepository.Get(services.Db, s.UserId)
		ret := repositories.UserRepository.IsFollows(services.Db, uid, s.UserId)
		info.IsFollow = ret.State
	}
	info.Id = u.Id
	info.Name = u.Name
	info.FollowerCount = u.FollowerCount
	info.FollowCount = u.FollowCount

	return info
}
