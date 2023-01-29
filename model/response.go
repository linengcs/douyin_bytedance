package model

type Response struct {
	StatusMsg  string `json:"status_msg"`
	StatusCode int    `json:"status_code"`
}

type UserInfo struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type VideoSimpleResponse struct {
	Model
	Title         string    `gorm:"not null" json:"title"`
	Author        *UserInfo `json:"author,omitempty"`
	PlayUrl       string    `json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count,omitempty"`
	CommentCount  int64     `json:"comment_count,omitempty"`
	IsFavorite    bool      `json:"is_favorite,omitempty"`
}

type CommentSimpleResponse struct {
	Model
	Author     *UserInfo `json:"author,omitempty"`
	Content    string    `json:"content,omitempty"`
	CreateDate string    `json:"create_date,omitempty"`
}

type VideoFeedResponse struct {
	Response
	VideoList []VideoSimpleResponse `json:"video_list"`
	NextTime  int                   `json:"next_time"`
}

type VideoListResponse struct {
	Response
	Videos []VideoSimpleResponse `json:"video_list"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserRegisterResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	Response
	UserInfo `json:"user"`
}

type VideoCommentResponse struct {
	Response
	CommentList []CommentSimpleResponse `json:"comment_list"`
}

type UserFollowResponse struct {
	Response
	UserList []UserInfo `json:"user_list"`
}
