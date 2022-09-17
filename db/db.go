package db

import (
	"fmt"

	"login-jwt/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SetupDB : initializing mysql database
func SetupDB() *gorm.DB {
	USER := config.GetConfig().DataBaseConfig.User
	PASS := config.GetConfig().DataBaseConfig.Pass
	HOST := config.GetConfig().DataBaseConfig.Host
	PORT := config.GetConfig().DataBaseConfig.Port
	DBNAME := config.GetConfig().DataBaseConfig.DB
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open(mysql.Open(URL), &gorm.Config{
		// hide log when insert data
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err.Error())
	}
	return db
}
