package user

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/forgocode/family/internal/pkg/newlog"
	"github.com/forgocode/family/internal/pkg/response"
	"github.com/forgocode/family/internal/pkg/sendlog"
	comemntService "github.com/forgocode/family/internal/webservice/service/comemnt"
	systemService "github.com/forgocode/family/internal/webservice/service/system"
	"github.com/forgocode/family/pkg/paginate"
)

func AdminGetUserInfo(ctx *gin.Context) {
	type info struct {
		UserID string `json:"userID" form:"userID"`
	}

	id := &info{}
	fmt.Println(id)
	err := ctx.ShouldBindQuery(id)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}

	type result struct {
		ArticleCount int64  `json:"articleCount"`
		NickName     string `json:"nickName"`
		Score        int64  `json:"score"`
		ShortComment int64  `json:"shortComment"`
		FollowCount  int64  `json:"followCount"`
		Description  string `json:"description"`
	}

	user, err := systemService.GetUserByUserID(id.UserID)
	if err != nil {
		return
	}
	shortCount, err := comemntService.GetCommentCountByUserID(id.UserID)
	if err != nil {
		return
	}
	r := &result{
		ArticleCount: 1,
		NickName:     user.NickName,
		Score:        user.Score,
		ShortComment: shortCount,
		FollowCount:  0,
		Description:  "蒸饭机器人",
	}
	response.Success(ctx, r, 1)

}

func AdminBanUser(ctx *gin.Context) {

}

func AdminUnBanUser(ctx *gin.Context) {

}

func AdminDeleteUser(ctx *gin.Context) {

}

func AdminCreateUser(ctx *gin.Context) {
	user := &systemService.UIUser{}
	err := ctx.ShouldBindJSON(user)
	if err != nil {
		newlog.Logger.Errorf("failed bind json, err: %+v\n", err)
		response.Failed(ctx, response.ErrStruct)
		return
	}
	err = systemService.AdminCreateUser(user)
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
	users, count, err := systemService.GetAllUser(q)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, users, count)
}
