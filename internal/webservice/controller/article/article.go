package article

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/forgocode/family/internal/pkg/response"
	articleService "github.com/forgocode/family/internal/webservice/service/article"
	"github.com/forgocode/family/pkg/paginate"
)

func CreateNewArticle(ctx *gin.Context) {
	info := &articleService.UIArticle{}
	err := ctx.ShouldBindJSON(info)
	if err != nil {
		fmt.Println(err)
		response.Failed(ctx, response.ErrStruct)
		return
	}
	info.AuthorID = ctx.Request.Header.Get("userID")
	info.UserName = ctx.Request.Header.Get("userName")
	err = articleService.CreateArticle(info)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, "", 1)
}

func AdminGetArticle(ctx *gin.Context) {
	q, err := paginate.GetPageQuery(ctx)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	articles, count, err := articleService.AdminGetArticleList(q)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, articles, count)
}

func NormalGetArticle(ctx *gin.Context) {
	q := &paginate.PageQuery{Page: 1, PageSize: 15}
	articles, err := articleService.NormalGetArticleList(q)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, articles, int64(len(articles)))
}

func NormalGetArticleInfo(ctx *gin.Context) {
	articleID := ctx.Param("id")
	info, err := articleService.GetArticleInfoByArticleID(articleID)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, info, 1)
}
