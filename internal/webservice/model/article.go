package model

type Article struct {
	AuthorID  string `json:"authorID" bson:"authorID" gorm:"column:authorID"`
	ArticleID string `json:"articleID" bson:"articleID" gorm:"column:articleID"`
	// 文章内容
	Context    string `json:"context" bson:"context" gorm:"column:context"`
	CreateTime int64  `json:"createTime" bson:"createTime" gorm:"column:createTime"`
	// 点赞数
	LikeCount      int32 `json:"likeCount" bson:"likeCount" gorm:"column:likeCount"`
	UnLikeCount    int32 `json:"unLikeCount" bson:"unLikeCount" gorm:"column:unLikeCount"`
	IsShow         int   `json:"isShow" bson:"isShow" gorm:"column:isShow"`
	IsShortArticle bool  `json:"isShortArticle" bson:"isShortArticle" gorm:"column:isShortArticle"`
}
