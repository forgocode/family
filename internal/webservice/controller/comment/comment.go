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
	com.AuthorID = "10000000"
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
