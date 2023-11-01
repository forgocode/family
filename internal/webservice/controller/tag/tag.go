package tag

import (
	"github.com/gin-gonic/gin"

	"github.com/forgocode/family/internal/webservice/service/tag"
)

func NormalGetAllTag(ctx *gin.Context) {

}

func AdminGetAllTag(ctx *gin.Context) {

}

func AdminUpdateTag(ctx *gin.Context) {
	tag.AdminUpdateTag("", true)
}

func AdminDeleteTag(ctx *gin.Context) {
	tag.AdminDeleteTag("")
}

func AdminCreateTag(ctx *gin.Context) {
	tag.AdminCreateTag(nil)
}
