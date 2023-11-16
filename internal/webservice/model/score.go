package model

// 记录每一个用户的每一个积分记录
type Score struct {
	UserID     string `json:"userID" gorm:"column:userID"`
	Score      int64  `json:"score" gorm:"column:score"`
	CreateTime int64  `json:"createTime" gorm:"column:createTime"`
	Reason     string `json:"reason" gorm:"column:reason"`
}
