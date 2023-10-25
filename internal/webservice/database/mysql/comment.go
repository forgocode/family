package mysql

import (
	"errors"

	"github.com/forgocode/family/internal/webservice/model"
)

func (c *MysqlClient) CreateComment(comment *model.CommunityComment) error {
	result := c.c.Create(comment)
	if result.Error != nil {
		return errors.New(result.Error.Error())
	}
	return nil
}

func (c *MysqlClient) BanComment(commentID string) error {
	currentCommentParent, err := c.getParentIDByCommentID(commentID)
	if err != nil {
		return err
	}
	currentCommentChild, err := c.getCommentIDByParentID(commentID)
	if err != nil {
		return err
	}

	//应当在事务中执行
	result := c.c.Model(&model.User{}).Where("commentID = ?", commentID).Update("isShow", false)
	if result.Error != nil {
		return nil
	}

	err = c.updateCommentParentID(currentCommentParent, currentCommentChild)
	if err != nil {
		return err
	}
	err = c.updateCommentParentID(currentCommentChild, currentCommentParent)
	if err != nil {
		return err
	}
	return nil

}

func (c *MysqlClient) UnBanComment(commentID string) error {
	result := c.c.Model(&model.User{}).Where("commentID = ?", commentID).Update("isShow", true)
	return result.Error
}

func (c *MysqlClient) updateCommentParentID(commentID, parentID string) error {
	return c.c.Model(model.CommunityComment{}).Where("commentID = ?", commentID).Update("parentID", parentID).Error
}

func (c *MysqlClient) getParentIDByCommentID(commentID string) (string, error) {
	var comment model.CommunityComment
	result := c.c.Where("commentID = ?", commentID).First(&comment)
	if result.Error != nil {
		return "", result.Error
	}
	return comment.ParentID, nil
}

func (c *MysqlClient) getCommentIDByParentID(parentID string) (string, error) {
	var comment model.CommunityComment
	result := c.c.Where("parentID = ?", parentID).First(&comment)
	if result.Error != nil {
		return "", result.Error
	}
	return comment.CommentID, nil
}
