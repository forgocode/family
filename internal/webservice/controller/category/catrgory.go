package category

import (
	"github.com/gin-gonic/gin"

	"github.com/forgocode/family/internal/pkg/newlog"
	"github.com/forgocode/family/internal/pkg/response"
	"github.com/forgocode/family/internal/webservice/service/category"
)

func NormalGetAllCategory(ctx *gin.Context) {

}

func AdminGetAllCategory(ctx *gin.Context) {

}

func AdminUpdateCategory(ctx *gin.Context) {
	category.AdminUpdateCategory("", true)
}

func AdminDeleteCategory(ctx *gin.Context) {
	category.AdminDeleteCategory("")

}

func AdminCreateCategory(ctx *gin.Context) {
	cate := &category.UICategory{}
	err := ctx.ShouldBindJSON(cate)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	newlog.Logger.Infof("%+v\n", cate)
	err = category.AdminCreateCategory(cate)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, "", 1)
}
