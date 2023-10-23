package model

type CommunityComment struct {
	AuthorID  string `json:"authorID" bson:"authorID" gorm:"column:authorID"`
	CommentID string `json:"commentID" bson:"commentID" gorm:"column:commentID"`
	// 评论内容
	Context    string `json:"context" bson:"context" gorm:"column:context"`
	CreateTime string `json:"createTime" bson:"createTime" gorm:"column:createTime"`
	// 点赞数
	LikeCount int64  `json:"likeCount" bson:"likeCount" gorm:"column:likeCount"`
	ParentID  string `json:"parentID" bson:"parentID" gorm:"column:parentID"`
}

func (CommunityComment) TableName() string {
	return "community_comment"
}
