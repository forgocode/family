package user

import (
	"github.com/gin-gonic/gin"

	"github.com/forgocode/family/internal/pkg/newlog"
	"github.com/forgocode/family/internal/pkg/response"
	"github.com/forgocode/family/internal/pkg/sendlog"
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

func AdminCreateUser(ctx *gin.Context) {
	user := &system.UIUser{}
	err := ctx.ShouldBindJSON(user)
	if err != nil {
		newlog.Logger.Errorf("failed bind json, err: %+v\n", err)
		response.Failed(ctx, response.ErrStruct)
		return
	}
	err = system.AdminCreateUser(user)
	if err != nil {
		newlog.Logger.Errorf("failed to create user: %+v, err: %+v\n", user, err)
		response.Failed(ctx, response.ErrDB)
		return
	}
	err = sendlog.SendOperationLog("root", "cn", sendlog.AddUser, user.NickName)
	if err != nil {
		newlog.Logger.Errorf("failed to send operation log: %+v, err: %+v\n", user, err)
	}
	response.Success(ctx, "", 1)
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
