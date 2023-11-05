package category

import (
	"github.com/gin-gonic/gin"

	"github.com/forgocode/family/internal/pkg/newlog"
	"github.com/forgocode/family/internal/pkg/response"
	"github.com/forgocode/family/internal/pkg/sendlog"
	"github.com/forgocode/family/internal/webservice/service/category"
	"github.com/forgocode/family/pkg/paginate"
)

func NormalGetAllCategory(ctx *gin.Context) {

}

func AdminGetAllCategory(ctx *gin.Context) {
	q, err := paginate.GetPageQuery(ctx)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	cates, count, err := category.AdminGetAllCategory(q)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, cates, count)

}

func AdminUpdateCategory(ctx *gin.Context) {
	category.AdminUpdateCategory("", true)
}

func AdminDeleteCategory(ctx *gin.Context) {
	uuid := ""
	err := ctx.ShouldBind(&uuid)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	err = category.AdminDeleteCategory(uuid)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	err = sendlog.SendOperationLog("root", "cn", sendlog.DeleteCategory, uuid)
	if err != nil {
		newlog.Logger.Errorf("failed to send operation log: %+v, err: %+v\n", uuid, err)
	}
	response.Success(ctx, "delete successfully", 1)

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
	err = sendlog.SendOperationLog("root", "cn", sendlog.NewCategory, cate.Name)
	if err != nil {
		newlog.Logger.Errorf("failed to send operation log: %+v, err: %+v\n", cate, err)
	}
	response.Success(ctx, "", 1)
}
