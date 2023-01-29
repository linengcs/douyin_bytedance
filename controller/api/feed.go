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

var FeedController = newFeedController()

type feedController struct {
	ctx *gin.Context
}

func newFeedController() *feedController {
	return &feedController{}
}

// Feed most 30 videos for every request
func (c *feedController) Feed(ctx *gin.Context) {
	latestTime := ctx.Query("latest_time")
	list := services.VideoService.GetVideos(latestTime)
	if len(list) == 0 {
		ctx.JSON(http.StatusOK, model.ErrFeedEmpty)
		return
	}
	var videos []model.VideoSimpleResponse

	token := ctx.Query("token")
	claims, _ := services.TokenService.ParseToken(token)
	if claims != nil {
		user := services.UserService.GetUserById(claims.UserID)
		var tag bool
		for i := 0; i < len(list); i++ {
			result, _ := services.RedisConn.Get("user_star" + strconv.Itoa(int(user.Id)) + ":" + strconv.Itoa(int(list[i].Id))).Result()
			if result != "1" {
				tag = false
			} else {
				tag = true
			}
			videos = append(videos, *render.BuildSimpleVideo(&list[i], tag, claims.UserID))
		}
	} else {
		for i := 0; i < len(list); i++ {
			videos = append(videos, *render.BuildSimpleVideoWithoutUser(&list[i], false))
		}
	}
	st := list[len(list)-1].CreateTime
	t, err := strconv.Atoi(st[:4] + st[5:7] + st[8:10] + st[11:13] + st[14:16] + st[17:19])
	if err != nil {
		log.Println(err)
	}
	ctx.JSON(http.StatusOK, model.VideoFeedResponse{
		Response: model.SUCCESS,
		VideoList: videos,
		NextTime:  t,
	},
	)
}
