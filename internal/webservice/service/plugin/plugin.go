package plugin

import (
	"github.com/forgocode/family/internal/webservice/database/mysql"
	"github.com/forgocode/family/internal/webservice/model"
	"github.com/forgocode/family/pkg/paginate"
)

func ListAllPlugin(q paginate.PageQuery) ([]model.Plugin, error) {
	c := mysql.GetClient()
	var plugins []model.Plugin
	result := c.C.Model(&model.Plugin{}).Scopes(paginate.ParseQuery(q)).Find(&plugins)
	return plugins, result.Error
}

func UpdatePluginStatusByName(status int, name string) error {
	c := mysql.GetClient()
	return c.C.Model(&model.Plugin{}).Where("name = ?", name).Update("status", status).Error
}

func DeletePluginByName(name string) error {
	c := mysql.GetClient()
	return c.C.Where("name = ?", name).Delete(&model.Plugin{}).Error
}

func CreatePlugin(p model.Plugin) error {
	c := mysql.GetClient()
	return c.C.Model(&model.Plugin{}).Create(&p).Error
}
