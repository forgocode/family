package category

import (
	"github.com/gin-gonic/gin"

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
	category.AdminCreateCategory(nil)
}
