package db

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Mysql       *gorm.DB
	Redis       *redis.Client
	Mongo       *mongo.Client
	dbConnector []Connector
)

type Connector interface {
	Connect()
}

func InitDb() {
	dbConnector = append(dbConnector, NewMysqlConnector())
	dbConnector = append(dbConnector, NewRedisConnector())
	dbConnector = append(dbConnector, NewMongoConnector())

	if len(dbConnector) == 0 {
		logrus.Warning("init db connector is empty!")
		return
	}

	for _, connector := range dbConnector {
		connector.Connect()
	}
}
