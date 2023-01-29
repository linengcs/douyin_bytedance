package main

import (
	"douyin/model"
	"douyin/services"
	"douyin/setting"
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	if err := setting.InitConfig(); err != nil{
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	services.ConnectDb()

	services.MqService.Consumer("", "go_star", Callback)
}


func Callback(s string){
	var msg = &model.FollowData{}
	err := json.Unmarshal([]byte(s), &msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = services.UserService.LikeVideo(msg.UserId, msg.VideoId, msg.State)
	if err != nil {
		fmt.Println(err)
	}

	redisConn := services.RedisService.Connect()
	defer redisConn.Close()
	redisConn.Set("user_star" + strconv.Itoa(int(msg.UserId)) + ":" + strconv.Itoa(int(msg.VideoId)), msg.State, 0 )
	fmt.Printf("msg is :%s\n", s)
}
