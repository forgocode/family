package tag

import (
	"time"

	"github.com/forgocode/family/internal/webservice/database/mysql"
	"github.com/forgocode/family/internal/webservice/model"
	"github.com/forgocode/family/pkg/uuid"
)

type UITag struct {
	Creator string `json:"creator"`
	Name    string `json:"name"`
	IsShow  bool   `json:"isShow"`
}

func (t *UITag) Convert() *model.Tag {
	return &model.Tag{
		CreateTime: time.Now().UnixMilli(),
		Creator:    t.Creator,
		Name:       t.Name,
		IsShow:     t.IsShow,
		UUID:       uuid.GetUUID(),
	}

}

func NormalGetAllTag() {
	c := mysql.GetClient()
	c.GetAllTag()
}

func AdminGetAllTag() {

}

func AdminCreateTag(tag *UITag) error {
	c := mysql.GetClient()

	return c.CreateTag(tag.Convert())
}

func AdminDeleteTag(uuid string) error {
	c := mysql.GetClient()
	return c.DeleteTag(uuid)
}

func AdminUpdateTag(uuid string, isShow bool) error {
	c := mysql.GetClient()
	if isShow {
		return c.UpdateTagShow(uuid)
	}
	return c.UpdateTagNotShow(uuid)
}
