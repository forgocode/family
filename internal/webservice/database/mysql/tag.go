package mysql

import (
	"github.com/forgocode/family/internal/webservice/model"
	"github.com/forgocode/family/pkg/paginate"
)

func (c *MysqlClient) CreateTag(cate *model.Tag) error {
	return c.c.Create(cate).Error
}

func (c *MysqlClient) UpdateTagShow(uuid string) error {
	return c.c.Where("uuid = ?", uuid).Update("isShow", true).Error
}

func (c *MysqlClient) UpdateTagNotShow(uuid string) error {
	return c.c.Where("uuid = ?", uuid).Update("isShow", false).Error
}

func (c *MysqlClient) GetAllTag(q *paginate.PageQuery) ([]model.Tag, error) {
	var tags []model.Tag
	result := c.c.Model(&model.Tag{}).Offset((q.Page - 1) * q.PageSize).Limit(q.PageSize).Find(&tags)
	return tags, result.Error

}
func (c *MysqlClient) GetTagCount() (int64, error) {
	var count int64
	result := c.c.Model(&model.Tag{}).Count(&count)
	return count, result.Error
}

func (c *MysqlClient) DeleteTag(uuid string) error {
	return c.c.Where("uuid = ?", uuid).Delete(&model.Tag{}).Error

}
