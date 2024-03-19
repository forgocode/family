package paginate

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PageQuery struct {
	//(page-1)*limit+1- page*limit
	Page       int         `json:"page" form:"page"`
	PageSize   int         `json:"pageSize" form:"pageSize"`
	StartTime  int32       `json:"startTime" form:"startTime"`
	EndTime    int32       `json:"endTime" form:"endTime"`
	Conditions []Condition `json:"conditions" form:"conditions"`
	Sorts      []Sort      `json:"sorts" from:"sorts"`
}

type Condition struct {
	Field     string      `json:"field" form:"field"`
	Value     interface{} `json:"value" form:"value"`
	Operation int         `json:"operation" form:"operation"`
}

type Sort struct {
	Field   string `json:"field" form:"field"`
	OrderBy int    `json:"orderBy" form:"orderBy"`
}

const (
	Equal        = iota + 1 // =
	NotEqual                // !=
	GreaterThan             //>
	GreaterEqual            //>=
	LessThan                //<
	LessEqual               //<=
	Like                    // like
	In
	NotIn
)

func GetPageQuery(ctx *gin.Context) (*PageQuery, error) {
	page := &PageQuery{}
	err := ctx.ShouldBindQuery(page)
	if err != nil {
		return nil, err
	}

	return page, nil
}

func Order(sorts []Sort) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		resultDB := db
		for _, sort := range sorts {
			if sort.OrderBy == 1 {
				resultDB.Order(sort.Field + " acs")
			}
			if sort.OrderBy == -1 {
				resultDB.Order(sort.Field + " desc")
			}
		}
		return resultDB
	}
}

func ParseQuery(q PageQuery) *gorm.DB {
	return nil
}

func StringFilter(key string, value interface{}, operation int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		resultDB := db
		if value == "" {
			return resultDB
		}
		switch operation {
		case Equal:
			resultDB = resultDB.Where(key+" = ?", value)
		case GreaterThan:
			resultDB = resultDB.Where(key+" > ?", value)
		case GreaterEqual:
			resultDB = resultDB.Where(key+" >= ?", value)
		case LessThan:
			resultDB = resultDB.Where(key+" < ?", value)
		case LessEqual:
			resultDB = resultDB.Where(key+" <= ?", value)
		case Like:
			if _, ok := value.(string); ok {
				resultDB = resultDB.Where(key+" Like ?", "%"+value.(string)+"%")
			}
		}
		return resultDB
	}
}

func ArrayFilter(key, string, value interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}
