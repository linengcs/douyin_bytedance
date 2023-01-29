package api

import (
	"douyin/controller/render"
	"douyin/model"
	"douyin/services"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var CommentController = newCommentService()

type commentController struct {
	ctx *gin.Context
}

func newCommentService() *commentController {
	return &commentController{}
}

func (c *commentController) CommentAction(ctx *gin.Context) {
	id, _ := ctx.Get("claim_id")
	uid, _ := id.(int64)

	vid, _ := strconv.ParseInt(ctx.Query("video_id"), 10, 64)
	action, _ := strconv.Atoi(ctx.Query("action_type"))
	ctxt := ctx.Query("comment_text")
	cid, _ := strconv.ParseInt(ctx.Query("comment_id"), 10, 64)

	if err := services.CommentService.CommentVideo(uid, vid, action, cid, ctxt); err != nil {
		if err.Error() == "no permission"{
			ctx.JSON(http.StatusOK, model.ErrCommentRight)
		}
		ctx.JSON(http.StatusOK, model.Failed)
	}
	ctx.JSON(http.StatusOK, model.SUCCESS)
}

func (c *commentController) CommentList(ctx *gin.Context) {
	id, _ := ctx.Get("claim_id")
	uid, _ := id.(int64)

	vid, _ := strconv.ParseInt(ctx.Query("video_id"), 10, 64)
	list := services.CommentService.GetComments(vid)
	var comments []model.CommentSimpleResponse
	for _, v := range list {
		comments = append(comments, *render.BuildSimpleComment(&v, uid))
	}
	ctx.JSON(http.StatusOK, model.VideoCommentResponse{
		Response: model.SUCCESS,
		CommentList: comments,
	})
}
