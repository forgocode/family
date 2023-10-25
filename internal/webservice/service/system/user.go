package system

import "github.com/forgocode/family/internal/webservice/database/mysql"

func GetUserByPhone(phone string, passwd string) (string, error) {
	c := mysql.GetClient()
	user, err := c.GetUserByPhone(phone, passwd)
	if err != nil {
		return "", err
	}

	return user.UserID, nil
}
