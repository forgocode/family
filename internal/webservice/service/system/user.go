package system

import (
	"errors"

	"github.com/forgocode/family/internal/webservice/database/mysql"
	"github.com/forgocode/family/internal/webservice/model"
	"github.com/forgocode/family/pkg/paginate"
)

func GetUserByPhone(phone string, passwd string) (string, error) {

	user, err := getUserByPhone(phone, passwd)
	if err != nil {
		return "", err
	}

	return user.UserID, nil
}

func GetAllUser(q *paginate.PageQuery) ([]model.User, int64, error) {

	users, err := getAllUser(q)
	if err != nil {
		return nil, 0, err
	}
	count, err := getUserCount()
	if err != nil {
		return nil, 0, err
	}
	return users, count, nil
}

func createUser(user *model.User) error {
	c := mysql.GetClient()
	result := c.C.Create(user)
	if result.Error != nil {
		return errors.New(result.Error.Error())
	}
	return nil
}

func isUserIDExist(userID string) (bool, error) {
	c := mysql.GetClient()
	var user model.User
	result := c.C.Where("userID = ?", userID).Find(&user)
	if result.Error != nil {
		return true, result.Error
	}
	if result.RowsAffected != 0 {
		return true, nil
	}
	return false, nil
}

func getUserCount() (int64, error) {
	c := mysql.GetClient()
	var count int64
	result := c.C.Model(&model.User{}).Count(&count)
	return count, result.Error
}

func getAllUser(q *paginate.PageQuery) ([]model.User, error) {
	c := mysql.GetClient()
	var users []model.User
	result := c.C.Model(&model.User{}).Offset((q.Page - 1) * q.PageSize).Limit(q.PageSize).Find(&users)
	return users, result.Error
}

func getUserByPhone(phone, passwd string) (*model.User, error) {
	c := mysql.GetClient()
	var user model.User
	result := c.C.Where("phone = ? AND password = ?", phone, passwd).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func getUserByUserID(userID string) (*model.User, error) {
	c := mysql.GetClient()
	var user model.User
	result := c.C.Where("userID = ?", userID).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func banUser(userID string) error {
	c := mysql.GetClient()
	result := c.C.Model(&model.User{}).Where("userID = ?", userID).Update("status", model.UserIsBaned)
	return result.Error
}

func unBanUser(userID string) error {
	c := mysql.GetClient()
	result := c.C.Model(&model.User{}).Where("userID = ?", userID).Update("status", model.UserIsNormal)
	return result.Error
}

func deleteUser(userID string) error {
	c := mysql.GetClient()
	result := c.C.Where("userID = ?", userID).Delete(&model.User{})
	return result.Error
}

//需要更新哪个就赋值哪个字段

func updateUserInfo(user model.User) error {
	c := mysql.GetClient()
	result := c.C.Where("userID = ?").Updates(user)
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
