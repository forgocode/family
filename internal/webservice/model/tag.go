package model

type Tag struct {
	CreateTime string `json:"createTime" gorm:"column:createTime"`
	Creator    string `json:"creator" gorm:"column:creator"`
	Name       string `json:"name" gorm:"column:name"`
	Status     string `json:"status" gorm:"column:status"`
	UUID       string `json:"uuid" gorm:"column:uuid"`
}

func (t Tag) TableName() string {
	return "tag"
}
