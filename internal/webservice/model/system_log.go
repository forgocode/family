package model

type SystemLog struct {
	CreateTime string `json:"createTime" bson:"createTime" gorm:"column:createTime"`
	Msg        string `json:"msg" bson:"msg" gorm:"column:msg"`
	Status     string `json:"status" bson:"status" gorm:"column:status"`
	Type       string `json:"type" bson:"type" gorm:"type"`
	UserID     string `json:"userID" bson:"userID" gorm:"userID"`
	UUID       string `json:"uuid" bson:"uuid" gorm:"uuid"`
}

func (SystemLog) TableName() string {
	return "system_log"
}
