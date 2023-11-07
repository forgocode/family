package mysql

import (
	"github.com/forgocode/family/internal/webservice/model"
	"github.com/forgocode/family/pkg/paginate"
)

func (c *MysqlClient) CreateCategory(cate *model.Category) error {
	return c.c.Create(cate).Error
}

func (c *MysqlClient) UpdateCategoryShow(uuid string) error {
	return c.c.Where("uuid = ?", uuid).Update("isShow", true).Error
}

func (c *MysqlClient) UpdateCategoryNotShow(uuid string) error {
	return c.c.Where("uuid = ?", uuid).Update("isShow", false).Error
}

func (c *MysqlClient) GetAllCategory(q *paginate.PageQuery) ([]model.Category, error) {
	var cates []model.Category
	result := c.c.Model(&model.Category{}).Order("createTime desc").Offset((q.Page - 1) * q.PageSize).Limit(q.PageSize).Find(&cates)
	return cates, result.Error

}
func (c *MysqlClient) GetCategoryCount() (int64, error) {
	var count int64
	result := c.c.Model(&model.Category{}).Count(&count)
	return count, result.Error
}

func (c *MysqlClient) DeleteCategory(uuid string) error {
	return c.c.Where("uuid = ?", uuid).Delete(&model.Category{}).Error

}
