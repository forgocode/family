package model

type Friend struct {
	UID        string `json:"uid" gorm:"column:uid"`
	FriendUID  string `json:"friendUID" gorm:"column:friendUID"`
	Createtime int64  `json:"createTime" gorm:"column:createTime"`
}
