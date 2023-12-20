package system

import (
	"github.com/gin-gonic/gin"

	"github.com/forgocode/family/internal/pkg/newlog"
	"github.com/forgocode/family/internal/pkg/response"
	"github.com/forgocode/family/internal/webservice/service/system"
	"github.com/forgocode/family/pkg/paginate"
)

func AdminGetOperationLog(ctx *gin.Context) {
	q, err := paginate.GetPageQuery(ctx)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	result, count, err := system.GetOperationLog(q)
	if err != nil {
		newlog.Logger.Errorf("error: %+v\n", err)
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, result, count)
}
