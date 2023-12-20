package article

import (
	"time"

	"github.com/forgocode/family/internal/webservice/database/mysql"
	"github.com/forgocode/family/internal/webservice/model"
	"github.com/forgocode/family/pkg/paginate"
	"github.com/forgocode/family/pkg/uuid"
)

type UIArticle struct {
	AuthorID string `json:"authorID"`
	UserName string `json:"userName"`

	ArticleID string `json:"articleID"`
	// 文章内容
	Context    string `json:"context" `
	CreateTime int64  `json:"createTime" `
	// 点赞数
	LikeCount    int32    `json:"likeCount"`
	IsOriginal   bool     `json:"isOriginal"`
	OriginalUrl  string   `json:"originalUrl"`
	OriginalUser string   `json:"originalUser"`
	IsShow       int      `json:"isShow"`
	ViewCount    int32    `json:"viewCount"`
	Introduction string   `json:"introduction"`
	Category     string   `json:"category"`
	Tags         []string `json:"tags"`
	Title        string   `json:"title"`
}

func (u *UIArticle) Convert() *model.Article {
	return &model.Article{
		AuthorID:     u.AuthorID,
		ArticleID:    uuid.GetUUID(),
		UserName:     u.UserName,
		Context:      u.Context,
		CreateTime:   time.Now().UnixMilli(),
		IsOriginal:   u.IsOriginal,
		OriginalUrl:  u.OriginalUrl,
		OriginalUser: u.OriginalUser,
		Title:        u.Title,
		Tags:         u.Tags,
		Category:     u.Category,
		Introduction: u.Introduction,
		IsShow:       u.IsShow,
	}
}

func CreateArticle(a *UIArticle) error {
	return createArticle(a.Convert())
}

func PublishArticle(articleID string) error {
	return publishArticle(articleID)
}

func BanArticle(articleID string) error {
	return banArticle(articleID)
}

func SendBackArticle(articleID string) error {
	return sendBackArticle(articleID)
}

func AdminGetArticleList(q *paginate.PageQuery) ([]model.Article, int64, error) {
	articles, err := getAllArticle(q)
	if err != nil {
		return nil, 0, err
	}
	count, err := getArticleCount()
	if err != nil {
		return nil, 0, err
	}
	return articles, count, nil
}

func NormalGetArticleList(q *paginate.PageQuery) ([]model.Article, error) {
	return normalGetAllArticle(q)
}

func GetArticleInfoByArticleID(id string) (model.Article, error) {
	return getArticleInfoByID(id)
}

func publishArticle(articleID string) error {
	c := mysql.GetClient()
	return c.C.Model(&model.Article{}).Where("articleID = ?", articleID).Update("isShow", model.ArticleShow).Error
}

func banArticle(articleID string) error {
	c := mysql.GetClient()
	return c.C.Model(&model.Article{}).Where("articleID = ?", articleID).Update("isShow", model.ArticleBanned).Error
}

func sendBackArticle(articleID string) error {
	c := mysql.GetClient()
	return c.C.Model(&model.Article{}).Where("articleID = ?", articleID).Update("isShow", model.ArticleSendBack).Error
}

func getAllArticle(q *paginate.PageQuery) ([]model.Article, error) {
	c := mysql.GetClient()
	var articles []model.Article
	result := c.C.Model(&model.Article{}).Where("isShow != ?", model.ArticleDraft).Order("createTime desc").Offset((q.Page - 1) * q.PageSize).Limit(q.PageSize).Find(&articles)
	return articles, result.Error
}

func normalGetAllArticle(q *paginate.PageQuery) ([]model.Article, error) {
	c := mysql.GetClient()
	var articles []model.Article
	result := c.C.Model(&model.Article{}).Where("isShow = ?", model.ArticleShow).Order("createTime desc").Offset((q.Page - 1) * q.PageSize).Limit(q.PageSize).Find(&articles)
	return articles, result.Error
}

func getArticleInfoByID(id string) (model.Article, error) {
	c := mysql.GetClient()
	var article model.Article
	result := c.C.Model(&model.Article{}).Where("articleID= ?", id).Find(&article)
	return article, result.Error
}

func createArticle(a *model.Article) error {
	c := mysql.GetClient()
	return c.C.Model(&model.Article{}).Create(a).Error
}

func getArticleCount() (int64, error) {
	c := mysql.GetClient()
	var count int64
	result := c.C.Model(&model.Article{}).Count(&count)
	return count, result.Error
}
