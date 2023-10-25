package database

import (
	"strconv"

	"github.com/forgocode/family/internal/pkg/newlog"
	"github.com/forgocode/family/internal/webservice/database/mysql"
	"github.com/forgocode/family/internal/webservice/model"
	"github.com/forgocode/family/pkg/userid"
)

func Start() {
	c, err := mysql.GetMysqlClient()
	if err != nil {
		return
	}
	err = c.AutoMigrate(
		&model.User{})
	if err != nil {
		newlog.Logger.Errorf("failed to auto migrate mysql table, err:%+v\n", err)
		return
	}
	createSuperAdminUser()

}

func createSuperAdminUser() {
	u := model.User{
		UserID:   strconv.Itoa(userid.SuperAdministrator),
		UserName: "超级管理员",
		Password: "123456",
		NickName: "超级管理员",
		Theme:    "",
		Avatar:   "",
		Role:     "超级管理员",
		RoleID:   "",
		Phone:    "13888888888",
		Email:    "forgocode@163.com",
		Status:   0,
	}
	c, err := mysql.GetMysqlClient()
	if err != nil {
		newlog.Logger.Errorf("failed to get mysql client, err:%+v\n", err)
		return
	}
	var user model.User
	result := c.Where("userID = ?", u.UserID).Find(&user)
	if result.Error != nil {
		newlog.Logger.Errorf("failed to get user by %+v, err:%+v\n", u, err)
		return
	}
	if result.RowsAffected == 0 {
		result = c.Create(&u)
		if result.Error != nil {
			newlog.Logger.Errorf("failed to create user info: %+v, err:%+v\n", u, err)
			return
		}
	}
}
