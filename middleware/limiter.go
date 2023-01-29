package middleware

import (
	"douyin/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

func LimiterMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		limiter := rate.NewLimiter(rate.Every(time.Second * 1), 10)
		ok := limiter.Allow()
		if ok != true {
			c.JSON(http.StatusOK, model.ErrLimiter)
			c.Abort()
			return
		}
	}
}
