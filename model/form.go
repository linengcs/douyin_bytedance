package model

type Action struct {
	UserId     int64  `form:"user_id"`
	Token      string `form:"token"`
	VideoId    string `form:"video_id"`
	ActionType string `form:"action_type"`
}

type CreateVideoForm struct {
	Name     string
	UserId   int64
	PlayUrl  string
	CoverUrl string
}
