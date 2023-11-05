package category

import (
	"time"

	"github.com/forgocode/family/internal/webservice/database/mysql"
	"github.com/forgocode/family/internal/webservice/model"
	"github.com/forgocode/family/pkg/paginate"
	"github.com/forgocode/family/pkg/uuid"
)

type UICategory struct {
	Creator string `json:"creator"`
	Name    string `json:"name"`
	IsShow  bool   `json:"isShow"`
}

func (t *UICategory) Convert() *model.Category {
	return &model.Category{
		CreateTime: time.Now().UnixMilli(),
		Creator:    t.Creator,
		Name:       t.Name,
		IsShow:     t.IsShow,
		UUID:       uuid.GetUUID(),
	}

}

func NormalGetAllCategory() {
	//c := mysql.GetClient()
	//c.GetAllCategory()
}

func AdminGetAllCategory(q *paginate.PageQuery) ([]model.Category, int64, error) {
	c := mysql.GetClient()
	cates, err := c.GetAllCategory(q)
	if err != nil {
		return nil, 0, err
	}
	count, err := c.GetCategoryCount()
	if err != nil {
		return nil, 0, err
	}
	return cates, count, nil
}

func AdminCreateCategory(Category *UICategory) error {
	c := mysql.GetClient()

	return c.CreateCategory(Category.Convert())
}

func AdminDeleteCategory(uuid string) error {
	c := mysql.GetClient()
	return c.DeleteCategory(uuid)
}

func AdminUpdateCategory(uuid string, isShow bool) error {
	c := mysql.GetClient()
	if isShow {
		return c.UpdateCategoryShow(uuid)
	}
	return c.UpdateCategoryNotShow(uuid)
}
