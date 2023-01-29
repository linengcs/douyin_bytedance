package api

import (
	"douyin/model"
	"douyin/services"
	"douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var LoginController = newLoginController()

type loginController struct {
	ctx *gin.Context
}

func newLoginController() *loginController {
	return &loginController{}
}

func (c *loginController) Register(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	if !util.CheckUsername(username){
		ctx.JSON(http.StatusOK, model.ErrUserNameFormat)
		return
	}

	if !util.CheckPassword(password){
		ctx.JSON(http.StatusOK, model.ErrPasswordFormat)
		return
	}

	s := services.UserService.CheckUsername(username)
	if s == true {
		ctx.JSON(http.StatusOK, model.ErrUserNameExist)
		return
	}
	user := services.UserService.SignUp(username, password)

	token, err := services.TokenService.GenToken(user.Id)
	if err != nil {
		ctx.JSON(http.StatusOK, model.ErrTokenSetUpFail)
		return
	}

	ctx.JSON(http.StatusOK, model.UserRegisterResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "register success",
		},
		UserId: user.Id,
		Token:  token,
	})
}

func (c *loginController) Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	if !util.CheckUsername(username){
		ctx.JSON(http.StatusOK, model.ErrUserNameFormat)
		return
	}

	if !util.CheckPassword(password){
		ctx.JSON(http.StatusOK, model.ErrPasswordFormat)
		return
	}

	user, err := services.UserService.SignIn(username, password)
	if err != nil {
		ctx.JSON(http.StatusOK, model.ErrPassWordWrong)
		return
	}
	token, err := services.TokenService.GenToken(user.Id)

	ctx.JSON(http.StatusOK, model.UserLoginResponse{
		Response: model.SUCCESS,
		UserId: user.Id,
		Token:  token,
	})
}

func (c *loginController) UserInfo(ctx *gin.Context) {
	cid, _ := ctx.Get("claim_id")
	uid, _ := cid.(int64)

	id, _ := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	userInfo := services.UserService.GetUserInfoById(uid, id)

	ctx.JSON(http.StatusOK, model.UserInfoResponse{
		Response: model.SUCCESS,
		UserInfo: *userInfo,
	})
}