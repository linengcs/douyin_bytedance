package api

import (
	"douyin/controller/render"
	"douyin/model"
	"douyin/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var VideoController = newVideoController()

type videoController struct {
	ctx *gin.Context
}

func newVideoController() *videoController {
	return &videoController{}
}

func (c *videoController) PostVideo(ctx *gin.Context) {
	cid, _ := ctx.Get("claim_id")
	uid, _ := cid.(int64)

	err := services.VideoService.Publish(ctx, uid)
	if err != nil {
		log.Println(err)
	}
	ctx.JSON(http.StatusOK, model.SUCCESS)
}

func (c *videoController) GetVideos(ctx *gin.Context) {
	cid, _ := ctx.Get("claim_id")
	uid, _ := cid.(int64)

	id, _ := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	list := services.VideoService.GetUserVideos(id)
	var videos []model.VideoSimpleResponse
	var tag bool
	for i := 0; i < len(list); i++ {
		result, _ := services.RedisConn.Get("user_star" + strconv.Itoa(int(id)) + ":" + strconv.Itoa(int(list[i].Id))).Result()
		if result != "1" {
			tag = false
		} else {
			tag = true
		}
		videos = append(videos, *render.BuildSimpleVideo(&list[i], tag, uid))
	}
	ctx.JSON(http.StatusOK, model.VideoListResponse{
		Response: model.SUCCESS,
		Videos: videos,
	},
	)
}
