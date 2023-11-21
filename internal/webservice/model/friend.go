package model

type Friend struct {
	UID        string `json:"uid" gorm:"column:uid"`
	FriendUID  string `json:"friendUID" gorm:"column:friendUID"`
	CreateTime int64  `json:"createTime" gorm:"column:createTime"`
}
