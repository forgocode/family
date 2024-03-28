package plugin

import (
	"github.com/forgocode/family/internal/pkg/response"
	pluginService "github.com/forgocode/family/internal/webservice/service/plugin"
	"github.com/forgocode/family/pkg/paginate"
	"github.com/gin-gonic/gin"
)

func AdminGetAllPlugin(ctx *gin.Context) {
	q, err := paginate.GetPageQuery(ctx)
	if err != nil {
		return
	}
	data, err := pluginService.ListAllPlugin(*q)
	if err != nil {
		return
	}
	response.Success(ctx, data, int64(len(data)))
}

func AdminDeletePlugin(ctx *gin.Context) {

}

func AdminCreatePlugin(ctx *gin.Context) {

}

func AdminUpdatePlugin(ctx *gin.Context) {
	type tmpT struct {
		Name   string `json:"name"`
		Status int    `json:"status"`
	}
	info := &tmpT{}
	err := ctx.ShouldBind(&info)
	if err != nil {
		return
	}
	if info.Name == "基础服务" {
		return
	}
	err = pluginService.UpdatePluginStatusByName(info.Status, info.Name)
	if err != nil {
		return
	}
	response.Success(ctx, "", 1)

}
