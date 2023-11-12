package system

import (
	"errors"
	"time"

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

func UpdateUserLastLogin(userID string) error {
	u := model.User{UserID: userID, LastLoginTime: time.Now().UnixMilli()}
	return updateUserInfo(u)
}

func UserAddTrend() (interface{}, error) {
	type resultInfo struct {
		Date  time.Time `json:"date"`
		Count int       `json:"count"`
	}
	type trendInfo struct {
		Date  string `json:"date"`
		Count int    `json:"count"`
	}
	c := mysql.GetClient()
	var info []resultInfo
	var trend []trendInfo

	result := c.C.Raw("select a.created_at as date,ifnull (b.count, 0) as count\nfrom(\n    SELECT curdate() as created_at\n    union all\n    SELECT date_sub(curdate(), interval 1 day) as created_at\n    union all\n    SELECT date_sub(curdate(), interval 2 day) as created_at\n    union all\n    SELECT date_sub(curdate(), interval 3 day) as created_at\n    union all\n    SELECT date_sub(curdate(), interval 4 day) as created_at\n    union all\n    SELECT date_sub(curdate(), interval 5 day) as created_at\n    union all\n    SELECT date_sub(curdate(), interval 6 day) as created_at\n) a left join (\nSELECT DATE_FORMAT( from_unixtime(createTime/1000),'%Y-%m-%d')  as date,count(*) as count FROM user GROUP BY date\n) b on a.created_at = b.date order by a.created_at asc;").Scan(&info)
	for _, i := range info {
		trend = append(trend, trendInfo{Date: i.Date.Format("2006/01/02"), Count: i.Count})
	}
	return trend, result.Error
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

// 需要更新哪个就赋值哪个字段
func updateUserInfo(user model.User) error {
	c := mysql.GetClient()
	result := c.C.Where("userID = ?", user.UserID).Updates(user)
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
	return oldobj

}
