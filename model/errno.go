package model

/**

给用户看到的错误码（以下简称错误码）包括轻推application层应用抛出的错误码，轻应用抛出的错误码及开放平台抛出的错误码。
错误码由[错误类型] [错误项]构成，共5位，前两位表示错误类型，后3位表示错误项。如[不存在的] [团队ID]。
规范只定义错误类型，错误项在各项目设计阶段进行补充。

**/

var SUCCESS = Response{StatusCode: 0, StatusMsg: "success"}
var Failed  = Response{StatusCode: -1, StatusMsg: "fail"}

var (
	// 限流器相关 101 开头
	ErrLimiter = Response{StatusCode: 10101, StatusMsg: "wait a minute"}

	// Token相关 102 开头
	//ErrTokenExpired   = Response{StatusCode: 10201, StatusMsg: "token expired"}
	ErrTokenSetUpFail = Response{StatusCode: 10202, StatusMsg: "token generate failed"}
	ErrTokenInvalid   = Response{StatusCode: 10204, StatusMsg: "token invalid"}

	// 视频相关 103 开头
	ErrFeedEmpty = Response{StatusCode: 10301, StatusMsg: "feed empty"}

	// 用户相关 104 开头
	ErrPassWordWrong  = Response{StatusCode: 10401, StatusMsg: "password wrong"}
	ErrUserNameFormat = Response{StatusCode: 10402, StatusMsg: "username format error"}
	ErrPasswordFormat = Response{StatusCode: 10402, StatusMsg: "password format error"}
	ErrUserNameExist  = Response{StatusCode: 10403, StatusMsg: "username exist"}

	// 评论相关 105 开头
	ErrCommentRight = Response{StatusCode: 10501, StatusMsg: "no permission"}

	// 点赞相关 106 开头
	ErrFavoriteRepeat = Response{StatusCode: 10601, StatusMsg: "favorite repeat"}
	ErrFavoriteMqFail = Response{StatusCode: 10602, StatusMsg: "favorite mq error"}

	// 关注、粉丝相关 110 开头
	ErrFollowRepeat   = Response{StatusCode: 11001, StatusMsg: "follow repeat"}
	ErrFollowYourself = Response{StatusCode: 11002, StatusMsg: "follow yourself"}
	ErrFollow         = Response{StatusCode: 11003, StatusMsg: "follow failed"}

)
