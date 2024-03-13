package group

import (
	"github.com/forgocode/family/internal/pkg/newlog"
	"github.com/forgocode/family/internal/pkg/response"
	group_service "github.com/forgocode/family/internal/webservice/service/group"
	"github.com/forgocode/family/pkg/paginate"
	"github.com/gin-gonic/gin"
)

func GetAllGroupByUserUID(ctx *gin.Context) {
	uid := ctx.Request.Header.Get("userID")
	result, err := group_service.GetAllGroupByUID(uid)
	if err != nil {
		newlog.Logger.Errorf("failed to get group by uid: %+s, err: %+v\n", uid, err)
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, result, int64(len(result)))
}

func GetMemberByGroupUID(ctx *gin.Context) {
	groupID := ctx.Param("id")
	page := paginate.PageQuery{
		Page:     0,
		PageSize: 20,
	}
	result, err := group_service.GetAllMemeberByGroupUID(groupID, page)
	if err != nil {
		newlog.Logger.Errorf("failed to get group member by : %+v, err: %+v\n", groupID, err)
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, result, int64(len(result)))
}
