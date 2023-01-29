package model

import "github.com/golang-jwt/jwt"

type Model struct {
	Id int64 `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
}

type MyClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

type User struct {
	Model
	Name          string `json:"name"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
}

type Video struct {
	Model
	Title         string `gorm:"not null" json:"title"`
	AuthorId      int64  `json:"author"`
	CreateTime    string `json:"create_time"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
}

type Comment struct {
	Model
	VideoId    int64
	UserId     int64
	Content    string
	CreateTime string
	State      bool
}

type UserFollow struct {
	UserId   int64
	FollowId int64
	State    bool
}

type UserVideo struct {
	Model
	UserId  int64
	VideoId int64
	State   bool
}

type FollowData struct {
	UserId  int64
	VideoId int64
	State   int
}