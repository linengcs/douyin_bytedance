package render

import "douyin/model"

func BuildSimpleComment(comment *model.Comment, cId int64) *model.CommentSimpleResponse {
	if comment == nil {
		return nil
	}

	rsp := &model.CommentSimpleResponse{}
	rsp.Id = comment.Id
	rsp.Author = BuildUserDefaultIfNull(comment.UserId, cId)
	rsp.Content = comment.Content
	rsp.CreateDate = comment.CreateTime[:4] +
		"-" + comment.CreateTime[5:7] +
		"-" + comment.CreateTime[8:10] +
		" " + comment.CreateTime[11:13] +
		":" + comment.CreateTime[14:16] +
		":" + comment.CreateTime[17:19]

	return rsp
}
