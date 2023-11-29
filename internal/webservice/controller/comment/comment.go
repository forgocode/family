package comment

import (
	"github.com/gin-gonic/gin"

	"github.com/forgocode/family/internal/pkg/response"
	commentService "github.com/forgocode/family/internal/webservice/service/comemnt"
	topicService "github.com/forgocode/family/internal/webservice/service/topic"
)

func UserCreateComment(ctx *gin.Context) {
	com := &commentService.UIComment{}
	err := ctx.ShouldBindJSON(com)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	if com.Context == "" {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	com.AuthorID = ctx.Request.Header.Get("userID")
	com.Address = ctx.Request.Header.Get("clientIP")
	com.User = ctx.Request.Header.Get("userName")
	err = commentService.UserCreateComment(com)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	err = topicService.AddUsedCountByTopicName(com.Topic)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, "successfully", 1)
}

func UserGetComment(ctx *gin.Context) {
	comments, err := commentService.UserGetComment()
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, comments, 1)
}

func UserGetFirstComment(ctx *gin.Context) {
	comments, err := commentService.UserGetFirstComment()
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, comments, 1)
}

func UserGetChildComment(ctx *gin.Context) {
	type info struct {
		CommentID string `json:"commentID" form:"commentID"`
	}
	var id info
	err := ctx.ShouldBindQuery(&id)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	result, err := commentService.UserGetCommentByCommentID(id.CommentID)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, result, 1)
}
