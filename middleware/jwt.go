package middleware

import (
	"douyin/model"
	"douyin/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}

		if token == "" {
			c.JSON(http.StatusUnauthorized, model.ErrTokenInvalid)
			c.Abort()
			return
		}
		claims, err := services.TokenService.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, model.ErrTokenInvalid)
			c.Abort()
			return
		}
		c.Set("claim_id", claims.UserID)
	}
}