package model

type User struct {
	UserID        string `json:"userID" gorm:"column:userID; index"`
	UserName      string `json:"userName" gorm:"column:userName;comment: the name of user"`
	Password      string `json:"-" gorm:"column:password; comment: the password of user"`
	NickName      string `json:"nickName" gorm:"column:nickName; comment: the nickName of user"`
	Theme         string `json:"theme" gorm:"column:theme; comment: the theme of user"`
	Avatar        string `json:"avatar" gorm:"column:avatar; comment: the avatar of user"`
	Role          string `json:"role" gorm:"column:role; comment: the role of user"`
	RoleID        string `json:"roleID" gorm:"column:roleID; comment: the roleID of user"`
	Phone         string `json:"phone" gorm:"column:phone; comment: the phone of user"`
	Email         string `json:"email" gorm:"column:email; comment: the email of user"`
	Status        int    `json:"status" gorm:"column:status; comment: the status of user"`
	CreateTime    int64  `bson:"createTime" gorm:"column:createTime"`
	LastLoginTime int64  `bson:"lastLoginTime" gorm:"column:lastLoginTime"`
}

const (
	UserIsNormal = 1
	UserIsBaned  = 2
)

func (User) TableName() string {
	return "user"
}
