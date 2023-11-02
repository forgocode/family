package paginate

import (
	"github.com/gin-gonic/gin"
)

type PageQuery struct {
	//(page-1)*limit+1- page*limit
	Page      int   `json:"page" form:"page"`
	PageSize  int   `json:"pageSize" form:"pageSize"`
	StartTime int32 `json:"startTime" form:"startTime"`
	EndTime   int32 `json:"endTime" from:"endTime"`
}

func GetPageQuery(ctx *gin.Context) (*PageQuery, error) {
	page := &PageQuery{}
	err := ctx.ShouldBindQuery(page)
	if err != nil {
		return nil, err
	}

	return page, nil
}
