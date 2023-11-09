package system

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/forgocode/family/internal/pkg/newlog"
	"github.com/forgocode/family/internal/pkg/response"
	"github.com/forgocode/family/internal/pkg/sendlog"
	"github.com/forgocode/family/internal/webservice/middleware"
	"github.com/forgocode/family/internal/webservice/model"
	"github.com/forgocode/family/internal/webservice/service/system"
)

type LoginInfo struct {
	Phone  string `json:"phone"`
	Passwd string `json:"password"`
}

func Login(ctx *gin.Context) {
	var info LoginInfo
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		response.Failed(ctx, response.ErrStruct, "struct error")
		return
	}
	if info.Phone == "" {
		response.Failed(ctx, response.ErrStruct, "please input phone")
		return
	}
	user := model.User{
		Phone:    info.Phone,
		Password: info.Passwd,
	}
	newlog.Logger.Infof("%+v\n", user)
	userID, err := system.GetUserByPhone(user.Phone, user.Password)
	if err != nil {
		//登陆失败
		response.Failed(ctx, response.ErrUserNameOrPassword)
		return
	}

	if userID == "" {
		response.Failed(ctx, response.ErrUserNameOrPassword, "user not exist")
		return
	}
	if system.UpdateUserLastLogin(userID) != nil {
		response.Failed(ctx, response.ErrDB, "")
		return
	}

	newlog.Logger.Infof("user <%s> login successfully\n", userID)
	token, err := middleware.GenerateToken(user.UserID, "", "")
	if err != nil {
		response.Failed(ctx, response.ErrUserNameOrPassword)
		return
	}
	err = middleware.StoreToken(token)
	if err != nil {
		response.Failed(ctx, response.ErrRedis)
		return
	}
	err = sendlog.SendOperationLog(userID, "cn", sendlog.LoginCode)
	if err != nil {
		newlog.Logger.Errorf("failed to send operation log, err: %+v\n", err)
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code":  200,
			"msg":   "handle successfully",
			"token": token,
			"role":  user.Role,
		})
}
