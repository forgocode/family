package mysql

import "github.com/forgocode/family/internal/webservice/model"

func (c *MysqlClient) CreateCategory(cate *model.Category) error {
	return c.c.Create(cate).Error
}

func (c *MysqlClient) UpdateCategoryShow(uuid string) error {
	return c.c.Where("uuid = ?", uuid).Update("isShow", true).Error
}

func (c *MysqlClient) UpdateCategoryNotShow(uuid string) error {
	return c.c.Where("uuid = ?", uuid).Update("isShow", false).Error
}

func (c *MysqlClient) GetAllCategory() {

}

func (c *MysqlClient) DeleteCategory(uuid string) error {
	return c.c.Where("uuid = ?", uuid).Delete(&model.Category{}).Error

}
