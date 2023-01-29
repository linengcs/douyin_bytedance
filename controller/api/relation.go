package api

import (
	"douyin/controller/render"
	"douyin/model"
	"douyin/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var RelationController = newRelationController()

type relationController struct {
	ctx *gin.Context
}

func newRelationController() *relationController {
	return &relationController{}
}

func (c *relationController) Follow(ctx *gin.Context) {
	var (
		actionType = ctx.Query("action_type")
		toUserId   = ctx.Query("to_user_id")
	)
	cid, _ := ctx.Get("claim_id")
	tid, _ := strconv.ParseInt(toUserId, 10, 64)
	action, _ := strconv.Atoi(actionType)

	result, err := services.RedisConn.Get("user_follow" + strconv.FormatInt(cid.(int64), 10) + ":" + strconv.Itoa(int(tid))).Result()
	if result == actionType {
		ctx.JSON(http.StatusOK, model.ErrFollowRepeat)
		return
	}

	if cid.(int64) == tid {
		ctx.JSON(http.StatusOK, model.ErrFollowYourself)
		return
	}

	status, err := services.UserService.FollowUser(cid.(int64), tid, action)
	if status == false || err != nil{
		ctx.JSON(http.StatusOK, model.ErrFollow)
		return
	}
	ctx.JSON(http.StatusOK, model.SUCCESS)
}

func (c *relationController) FollowList(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	info := services.UserService.GetFollows(id)
	var follows []model.UserInfo
	for _, u := range info {
		follows = append(follows, *render.BuildUserInfo(&u, "follow", id))
	}

	ctx.JSON(http.StatusOK, model.UserFollowResponse{
		Response: model.SUCCESS,
		UserList: follows,
	},
	)
}

func (c *relationController) FollowerList(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	list := services.UserService.GetFollowers(id)
	var followers []model.UserInfo
	for _, u := range list {
		followers = append(followers, *render.BuildUserInfo(&u, "follower", id))
	}
	ctx.JSON(http.StatusOK, model.UserFollowResponse{
		Response: model.SUCCESS,
		UserList: followers,
	},
	)
}
