package render

import (
	"douyin/model"
	"douyin/repositories"
	"douyin/services"
)

func BuildSimpleVideo(video *model.Video, status bool, uId int64) *model.VideoSimpleResponse {
	if video == nil {
		return nil
	}

	rsp := &model.VideoSimpleResponse{}
	rsp.Id = video.Id
	rsp.Title = video.Title
	rsp.Author = BuildUserDefaultIfNull(video.AuthorId, uId)
	rsp.CommentCount = video.CommentCount
	rsp.CoverUrl = video.CoverUrl
	rsp.FavoriteCount = video.FavoriteCount
	rsp.IsFavorite = status
	rsp.PlayUrl = video.PlayUrl

	return rsp
}

func BuildSimpleVideoWithoutUser(video *model.Video, status bool) *model.VideoSimpleResponse {
	if video == nil {
		return nil
	}

	rsp := &model.VideoSimpleResponse{}
	rsp.Id = video.Id
	rsp.Title = video.Title
	rsp.Author = repositories.UserRepository.GetUserInfoById(services.Db, video.AuthorId)
	rsp.CommentCount = video.CommentCount
	rsp.CoverUrl = video.CoverUrl
	rsp.FavoriteCount = video.FavoriteCount
	rsp.IsFavorite = status
	rsp.PlayUrl = video.PlayUrl

	return rsp
}