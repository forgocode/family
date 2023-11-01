package model

type Tag struct {
	CreateTime int64  `json:"createTime" gorm:"column:createTime"`
	Creator    string `json:"creator" gorm:"column:creator"`
	Name       string `json:"name" gorm:"column:name"`
	IsShow     bool   `json:"isShow" gorm:"column:isShow"`
	UUID       string `json:"uuid" gorm:"column:uuid"`
}

func (t Tag) TableName() string {
	return "tag"
}
