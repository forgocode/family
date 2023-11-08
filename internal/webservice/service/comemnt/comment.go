package comemnt

import "C"
import (
	"errors"
	"time"

	"github.com/forgocode/family/internal/webservice/database/mysql"
	"github.com/forgocode/family/internal/webservice/model"
	"github.com/forgocode/family/pkg/uuid"
)

type UIComment struct {
	User         string `json:"user"`
	Context      string `json:"context"`
	AuthorID     string `json:"authorID" `
	CommentID    string `json:"commentID"`
	TopCommentID string `json:"topCommentID"`
	CreateTime   int64  `json:"createTime" `
	// 点赞数
	LikeCount   int32       `json:"likeCount" `
	UnLikeCount int32       `json:"unLikeCount" `
	ParentID    string      `json:"parentID" `
	IsShow      bool        `json:"isShow" `
	IsFirst     bool        `json:"isFirst"`
	Child       []UIComment `json:"child"`
	ReplayTo    string      `json:"replayTo"`
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

func UserGetComment() ([]UIComment, error) {
	var comments []UIComment
	first, err := findFirstLevelComment()
	if err != nil {
		return comments, err
	}
	for _, fc := range first {
		f := UIComment{
			User:        fc.AuthorID,
			Context:     fc.Context,
			AuthorID:    fc.AuthorID,
			CommentID:   fc.CommentID,
			CreateTime:  fc.CreateTime,
			LikeCount:   fc.LikeCount,
			UnLikeCount: fc.UnLikeCount,
			ParentID:    fc.ParentID,
			IsShow:      fc.IsShow,
			IsFirst:     fc.IsFirst,
			Child:       nil,
		}
		second, err := findSecondLevelComment(fc.CommentID, fc.CommentID)
		if err != nil {
			return comments, err
		}
		for _, sc := range second {
			s := UIComment{
				User:        sc.AuthorID,
				Context:     sc.Context,
				AuthorID:    sc.AuthorID,
				CommentID:   sc.CommentID,
				CreateTime:  sc.CreateTime,
				LikeCount:   sc.LikeCount,
				UnLikeCount: sc.UnLikeCount,
				ParentID:    sc.ParentID,
				IsShow:      sc.IsShow,
				IsFirst:     sc.IsFirst,
				Child:       nil,
			}
			if sc.ParentID == "" {
				continue
			}
			third, err := findThirdLevelComment(sc.TopCommentID, sc.CommentID, sc.AuthorID)
			if err != nil {
				return comments, err
			}
			for _, v := range third {
				t := UIComment{
					User:        v.AuthorID,
					Context:     v.Context,
					AuthorID:    v.AuthorID,
					CommentID:   v.CommentID,
					CreateTime:  v.CreateTime,
					LikeCount:   v.LikeCount,
					UnLikeCount: v.UnLikeCount,
					ParentID:    v.ParentID,
					IsShow:      v.IsShow,
					IsFirst:     v.IsFirst,
					Child:       nil,
				}
				s.Child = append(s.Child, t)
			}
			f.Child = append(f.Child, s)
		}
		comments = append(comments, f)
	}
	return comments, err
}

func createComment(comment *model.CommunityComment) error {
	c := mysql.GetClient()
	result := c.C.Create(comment)
	if result.Error != nil {
		return errors.New(result.Error.Error())
	}
	return nil
}

func findFirstLevelComment() ([]model.CommunityComment, error) {
	c := mysql.GetClient()
	var comments []model.CommunityComment
	result := c.C.Model(&model.CommunityComment{}).Where("parentID = ''").Order("createTime desc").Offset(0).Limit(10).Find(&comments)
	return comments, result.Error
}

func findSecondLevelComment(topCommentID, parentID string) ([]model.CommunityComment, error) {
	c := mysql.GetClient()
	var comments []model.CommunityComment
	result := c.C.Model(&model.CommunityComment{}).Where("topCommentID = ? AND parentID = ?", topCommentID, parentID).Order("createTime desc").Offset(0).Limit(10).Find(&comments)
	return comments, result.Error
}

func findThirdLevelComment(topCommentID, parentID string, replayTo string) ([]model.CommunityComment, error) {
	c := mysql.GetClient()
	var comments []model.CommunityComment
	result := c.C.Model(&model.CommunityComment{}).Where("topCommentID = ? AND parentID = ? AND replayTo = ?", topCommentID, parentID, replayTo).Order("createTime desc").Offset(0).Limit(10).Find(&comments)
	return comments, result.Error
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
