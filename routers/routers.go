package routers

import (
	"douyin/controller/api"
	"douyin/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")
	{

		apiRouter.GET("/feed/", api.FeedController.Feed)

		apiRouter.POST("/user/register/", api.LoginController.Register)
		apiRouter.POST("/user/login/", api.LoginController.Login)
		apiRouter.GET("/user/", middleware.JWTAuthMiddleware(), api.LoginController.UserInfo)

		publish := apiRouter.Group("/publish", middleware.JWTAuthMiddleware())
		{
			publish.POST("/action/", api.VideoController.PostVideo)
			apiRouter.GET("/list/", api.VideoController.GetVideos)
		}

		favorite := apiRouter.Group("/favorite")
		{
			favorite.POST("/action/", middleware.LimiterMiddleware(), middleware.JWTAuthMiddleware(), api.FavoriteController.Favorite)
			favorite.GET("/list/", middleware.JWTAuthMiddleware(), api.FavoriteController.GetFavorite)
		}

		comment := apiRouter.Group("/comment", middleware.JWTAuthMiddleware())
		{
			comment.POST("/action/", api.CommentController.CommentAction)
			comment.GET("/list/", api.CommentController.CommentList)
		}

		relation := apiRouter.Group("/relation")
		{
			relation.POST("/action/", middleware.LimiterMiddleware(), middleware.JWTAuthMiddleware(), api.RelationController.Follow)
			relation.GET("/follow/list/", middleware.JWTAuthMiddleware(), api.RelationController.FollowList)
			relation.GET("/follower/list/", middleware.JWTAuthMiddleware(), api.RelationController.FollowerList)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status_code": "404",
		})
	})
}
