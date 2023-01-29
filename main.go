package main

import (
	"douyin/middleware"
	"douyin/routers"
	"douyin/services"
	"douyin/setting"
	"douyin/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := setting.InitConfig(); err != nil{
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	services.ConnectDb()

	if err := middleware.InitLogger(setting.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
		return
	}

	if err := util.GetBuck(); err != nil {
		fmt.Printf("bucket get failed, err:%v\n", err)
		return
	}

	services.RedisConn = services.RedisService.Connect()
	defer services.RedisConn.Close()

	r := gin.Default()
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	routers.InitRouter(r)
	_ = r.Run(":8081")
}