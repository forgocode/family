package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/forgocode/family/internal/conf"
)

func GetMongoClient() (*mongo.Client, error) {
	if c == nil {
		config := conf.GetConfig()
		return InitMongo(config.Mongo.User, config.Mongo.Password, config.Mongo.IP, config.Mongo.Port, config.Mongo.DB)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := c.Ping(ctx, readpref.Primary())
	if err != nil {
		config := conf.GetConfig()
		c, err = InitMongo(config.Mongo.User, config.Mongo.Password, config.Mongo.IP, config.Mongo.Port, config.Mongo.DB)
		if err != nil {
			return nil, err
		}

	}
	return c, nil

}

var c *mongo.Client

func InitMongo(user string, passwd string, ip string, port uint16, db string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", user, passwd, ip, port, db)).SetConnectTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}
	return client, nil
}
