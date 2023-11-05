package user

import (
	"github.com/gin-gonic/gin"

	"github.com/forgocode/family/internal/pkg/response"
	"github.com/forgocode/family/internal/webservice/service/system"
	"github.com/forgocode/family/pkg/paginate"
)

func AdminGetUserInfo(ctx *gin.Context) {

}

func AdminBanUser(ctx *gin.Context) {

}

func AdminUnBanUser(ctx *gin.Context) {

}

func AdminDeleteUser(ctx *gin.Context) {

}

func NormalGetAllUser(ctx *gin.Context) {

	q, err := paginate.GetPageQuery(ctx)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	users, count, err := system.GetAllUser(q)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, users, count)
}
