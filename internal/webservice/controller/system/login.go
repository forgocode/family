package system

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/forgocode/family/internal/pkg/newlog"
	"github.com/forgocode/family/internal/pkg/response"
	"github.com/forgocode/family/internal/pkg/sendlog"
	"github.com/forgocode/family/internal/webservice/middleware"
	"github.com/forgocode/family/internal/webservice/model"
)

type LoginInfo struct {
	Phone  string `json:"phone"`
	Passwd string `json:"passwd"`
}

func Login(ctx *gin.Context) {
	err := sendlog.SendOperationLog("10001", "en", sendlog.LoginCode)
	if err != nil {
		return
	}
	var info LoginInfo
	err = ctx.ShouldBindJSON(&info)
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
	userName, ok := "", true
	if userName == "" {
		//response.Failed(ctx, response.ErrUserNameOrPassword)
		return
	}
	ctx.Request.Header.Set("userName", userName)
	if !ok {
		//service.SendOperationLog(ctx, "Login", "登陆失败")
		response.Failed(ctx, response.ErrUserNameOrPassword)
		return
	}
	newlog.Logger.Infof("user <%s> login successfully\n", userName)
	token, err := middleware.GenerateToken(user.Phone, userName)
	if err != nil {
		response.Failed(ctx, response.ErrUserNameOrPassword)
		return
	}
	err = middleware.StoreToken(token)
	if err != nil {
		response.Failed(ctx, response.ErrRedis)
		return
	}
	//service.SendOperationLog(ctx, "Login", "登陆成功")
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code":  200,
			"msg":   "handle successfully",
			"token": token,
		})
}
