package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/forgocode/family/internal/conf"
)

var c *sql.DB

func GetMysqlClient() (*sql.DB, error) {
	if c == nil {
		config := conf.GetConfig()
		return InitMySql(config.Mysql.IP, config.Mysql.User, config.Mysql.Password, config.Mysql.DB, config.Mysql.Port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := c.PingContext(ctx)
	if err != nil {
		config := conf.GetConfig()
		c, err = InitMySql(config.Mysql.IP, config.Mysql.User, config.Mysql.Password, config.Mysql.DB, config.Mysql.Port)
		if err != nil {
			return nil, err
		}

	}
	return c, nil

}

func InitMySql(url, user, passwd, dbName string, port uint16) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, passwd, url, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	client, err := db.DB()
	if err != nil {
		return nil, err
	}
	return client, nil
}
