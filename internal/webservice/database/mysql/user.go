package mysql

import (
	"errors"

	"github.com/forgocode/family/internal/webservice/model"
	"github.com/forgocode/family/pkg/paginate"
)

func (c *MysqlClient) CreateUser(user *model.User) error {
	result := c.c.Create(user)
	if result.Error != nil {
		return errors.New(result.Error.Error())
	}
	return nil
}

func (c *MysqlClient) IsUserIDExist(userID string) (bool, error) {
	var user model.User
	result := c.c.Where("userID = ?", userID).Find(&user)
	if result.Error != nil {
		return true, result.Error
	}
	if result.RowsAffected != 0 {
		return true, nil
	}
	return false, nil
}

func (c *MysqlClient) GetUserCount() (int64, error) {
	var count int64
	result := c.c.Model(&model.User{}).Count(&count)
	return count, result.Error
}

func (c *MysqlClient) GetAllUser(q *paginate.PageQuery) ([]model.User, error) {
	var users []model.User
	result := c.c.Model(&model.User{}).Offset((q.Page - 1) * q.PageSize).Limit(q.PageSize).Find(&users)
	return users, result.Error
}

func (c *MysqlClient) GetUserByPhone(phone, passwd string) (*model.User, error) {
	var user model.User
	result := c.c.Where("phone = ? AND password = ?", phone, passwd).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (c *MysqlClient) GetUserByUserID(userID string) (*model.User, error) {
	var user model.User
	result := c.c.Where("userID = ?", userID).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (c *MysqlClient) BanUser(userID string) error {
	result := c.c.Model(&model.User{}).Where("userID = ?", userID).Update("status", model.UserIsBaned)
	return result.Error
}

func (c *MysqlClient) UnBanUser(userID string) error {
	result := c.c.Model(&model.User{}).Where("userID = ?", userID).Update("status", model.UserIsNormal)
	return result.Error
}

func (c *MysqlClient) DeleteUser(userID string) error {
	result := c.c.Where("userID = ?", userID).Delete(&model.User{})
	return result.Error
}

//需要更新哪个就赋值哪个字段

func (c *MysqlClient) UpdateUserInfo(user model.User) error {

	return c.updateInfo(user)

}

func (c *MysqlClient) updateInfo(user model.User) error {
	result := c.c.Where("userID = ?").Updates(user)
	return result.Error
}

func getNewUserInfo(newobj, oldobj model.User) model.User {
	if newobj.NickName != "" {
		oldobj.NickName = newobj.NickName
	}
	if newobj.Email != "" {
		oldobj.Email = newobj.Email
	}
	if newobj.Password != "" {
		oldobj.Password = newobj.Password
	}
	if newobj.Phone != "" {
		oldobj.Phone = newobj.Phone
	}
	if newobj.Avatar != "" {
		oldobj.Avatar = newobj.Avatar
	}
	if newobj.Role != "" {
		oldobj.Role = newobj.Role
	}
	return oldobj

}
