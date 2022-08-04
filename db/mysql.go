package db

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGorm() *gorm.DB {

	connection := viper.GetString("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}
