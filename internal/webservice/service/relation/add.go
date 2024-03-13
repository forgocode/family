package relation

import (
	"time"

	"github.com/forgocode/family/internal/webservice/database/mysql"
	"github.com/forgocode/family/internal/webservice/model"
)

func AddUserGroupRelation(userUID, groupUID string, relation int) error {
	c := mysql.GetClient()
	r := &model.GroupUserRelation{
		GroupUID:   groupUID,
		UserUID:    userUID,
		Relation:   relation,
		CreateTime: time.Now().UnixMilli(),
	}
	result := c.C.Model(&model.GroupUserRelation{}).Create(r)
	return result.Error
}
