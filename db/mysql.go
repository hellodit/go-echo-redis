package db

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMysql() *gorm.DB {
	var dbHost = viper.GetString("DB_HOST")
	var dbPort = viper.GetString("DB_PORT")
	var dbUser = viper.GetString("DB_USER")
	var dbPass = viper.GetString("DB_PASSWORD")
	var dbName = viper.GetString("DB_NAME")

	parse := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(parse), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}
