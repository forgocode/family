package category

import (
	"time"

	"github.com/forgocode/family/internal/webservice/database/mysql"
	"github.com/forgocode/family/internal/webservice/model"
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
	c := mysql.GetClient()
	c.GetAllCategory()
}

func AdminGetAllCategory() {

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
