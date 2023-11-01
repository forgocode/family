package mysql

import "github.com/forgocode/family/internal/webservice/model"

func (c *MysqlClient) CreateTag(cate *model.Tag) error {
	return c.c.Create(cate).Error
}

func (c *MysqlClient) UpdateTagShow(uuid string) error {
	return c.c.Where("uuid = ?", uuid).Update("isShow", true).Error
}

func (c *MysqlClient) UpdateTagNotShow(uuid string) error {
	return c.c.Where("uuid = ?", uuid).Update("isShow", false).Error
}

func (c *MysqlClient) GetAllTag() ([]model.Tag, error) {
	return nil, nil

}

func (c *MysqlClient) DeleteTag(uuid string) error {
	return c.c.Where("uuid = ?", uuid).Delete(&model.Tag{}).Error

}
