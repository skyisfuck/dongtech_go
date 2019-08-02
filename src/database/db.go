package database

import (
	"dongtech_go/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"time"
)

func GetDB() *gorm.DB {
	config, err := config.GetConfig()
	if err != nil {
		logrus.Println("get config err")
		return nil
	}
	instance, err := gorm.Open("postgres", connStr(config))
	if err != nil {
		return nil
	}
	instance.DB().SetConnMaxLifetime(time.Minute * 5)
	instance.DB().SetMaxIdleConns(10)
	instance.DB().SetMaxOpenConns(config.Database.PoolSize)
	instance.LogMode(true)
	return instance
}

func connStr(config *config.Config) string {
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%d",
		config.Database.Addr, config.Database.User, config.Database.Database, config.Database.Password, config.Database.Port)
	return connStr
}
