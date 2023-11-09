package model

type GiveLike struct {
	ID       string `json:"ID" gorm:"column:ID"`
	UserID   string `json:"UserID" gorm:"column:userID"`
	UserName string `json:"userName" gorm:"column:userName"`
	Like     bool   `json:"like" gorm:"column:like"`
	Type     int8   `json:"type" gorm:"column:type"`
}

func (GiveLike) TableName() string {
	return "give_name"
}
