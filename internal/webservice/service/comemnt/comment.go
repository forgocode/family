package comemnt

import (
	"errors"
	"time"

	"github.com/forgocode/family/internal/webservice/database/mysql"
	"github.com/forgocode/family/internal/webservice/model"
	"github.com/forgocode/family/pkg/uuid"
)

type UIComment struct {
	User       string `json:"user"`
	Context    string `json:"context"`
	AuthorID   string `json:"authorID" `
	CommentID  string `json:"commentID"`
	CreateTime int64  `json:"createTime" `
	// 点赞数
	LikeCount   int32  `json:"likeCount" `
	UnLikeCount int32  `json:"unLikeCount" `
	ParentID    string `json:"parentID" `
	IsShow      bool   `json:"isShow" `
	IsFirst     bool   `json:"isFirst"`
}

func (c *UIComment) Convert() *model.CommunityComment {
	return &model.CommunityComment{
		AuthorID:    c.AuthorID,
		CommentID:   uuid.GetUUID(),
		Context:     c.Context,
		CreateTime:  time.Now().UnixMilli(),
		LikeCount:   0,
		UnLikeCount: 0,
		ParentID:    c.ParentID,
		IsShow:      true,
		IsFirst:     false,
	}
}

func UserCreateComment(comment *UIComment) error {
	return createComment(comment.Convert())
}

func createComment(comment *model.CommunityComment) error {
	c := mysql.GetClient()
	result := c.C.Create(comment)
	if result.Error != nil {
		return errors.New(result.Error.Error())
	}
	return nil
}

func banComment(commentID string) error {
	c := mysql.GetClient()
	currentCommentParent, err := getParentIDByCommentID(commentID)
	if err != nil {
		return err
	}
	currentCommentChild, err := getCommentIDByParentID(commentID)
	if err != nil {
		return err
	}

	//应当在事务中执行
	result := c.C.Model(&model.User{}).Where("commentID = ?", commentID).Update("isShow", false)
	if result.Error != nil {
		return nil
	}

	err = updateCommentParentID(currentCommentParent, currentCommentChild)
	if err != nil {
		return err
	}
	err = updateCommentParentID(currentCommentChild, currentCommentParent)
	if err != nil {
		return err
	}
	return nil

}

func unBanComment(commentID string) error {
	c := mysql.GetClient()
	result := c.C.Model(&model.User{}).Where("commentID = ?", commentID).Update("isShow", true)
	return result.Error
}

func updateCommentParentID(commentID, parentID string) error {
	c := mysql.GetClient()
	return c.C.Model(model.CommunityComment{}).Where("commentID = ?", commentID).Update("parentID", parentID).Error
}

func getParentIDByCommentID(commentID string) (string, error) {
	c := mysql.GetClient()
	var comment model.CommunityComment
	result := c.C.Where("commentID = ?", commentID).First(&comment)
	if result.Error != nil {
		return "", result.Error
	}
	return comment.ParentID, nil
}

func getCommentIDByParentID(parentID string) (string, error) {
	c := mysql.GetClient()
	var comment model.CommunityComment
	result := c.C.Where("parentID = ?", parentID).First(&comment)
	if result.Error != nil {
		return "", result.Error
	}
	return comment.CommentID, nil
}
