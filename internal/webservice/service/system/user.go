package system

import (
	"github.com/forgocode/family/internal/webservice/database/mysql"
	"github.com/forgocode/family/internal/webservice/model"
	"github.com/forgocode/family/pkg/paginate"
)

func GetUserByPhone(phone string, passwd string) (string, error) {
	c := mysql.GetClient()
	user, err := c.GetUserByPhone(phone, passwd)
	if err != nil {
		return "", err
	}

	return user.UserID, nil
}

func GetAllUser(q *paginate.PageQuery) ([]model.User, int64, error) {
	c := mysql.GetClient()
	users, err := c.GetAllUser(q)
	if err != nil {
		return nil, 0, err
	}
	count, err := c.GetUserCount()
	if err != nil {
		return nil, 0, err
	}
	return users, count, nil
}
