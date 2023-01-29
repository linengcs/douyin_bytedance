package api

import (
	"douyin/controller/render"
	"douyin/model"
	"douyin/services"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var FavoriteController = newFavoriteController()

type favoriteController struct {
	ctx *gin.Context
}

func newFavoriteController() *favoriteController {
	return &favoriteController{}
}

func (c *favoriteController) Favorite(ctx *gin.Context) {
	id, _ := ctx.Get("claim_id")
	uid, _ := id.(int64)

	actionType := ctx.Query("action_type")
	status, _ := strconv.Atoi(actionType)
	videoId, _ := strconv.ParseInt(ctx.Query("video_id"), 10, 64)

	result, err := services.RedisConn.Get("user_star" + strconv.FormatInt(uid, 10) + ":" + strconv.Itoa(int(videoId))).Result()
	if result == actionType {
		ctx.JSON(http.StatusOK, model.ErrFavoriteRepeat)
		return
	}

	var msg = &model.FollowData{
		UserId:  uid,
		VideoId: videoId,
		State:   status,
	}
	msgStr, _ := json.Marshal(msg)

	err = services.MqService.Publish("", "go_star", string(msgStr))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, model.ErrFavoriteMqFail)
		return
	}
	ctx.JSON(http.StatusOK, model.SUCCESS)
}

func (c *favoriteController) GetFavorite(ctx *gin.Context) {
	id, _ := ctx.Get("claim_id")
	uid, _ := id.(int64)

	list := services.VideoService.GetLikes(uid)

	var videos []model.VideoSimpleResponse
	for _, v := range list {
		videos = append(videos, *render.BuildSimpleVideo(&v, true, uid))
	}
	ctx.JSON(http.StatusOK, model.VideoListResponse{
		Response: model.SUCCESS,
		Videos: videos,
	},
	)
}
