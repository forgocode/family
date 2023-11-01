package paginate

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type PageQuery struct {
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	StartTime int32 `json:"startTime"`
	EndTime   int32 `json:"endTime"`
}

func GetPageQuery(ctx *gin.Context) (*PageQuery, error) {
	page := &PageQuery{}
	s, ok := ctx.GetQuery("page")
	fmt.Println(s, ok)

	return page, nil
}
