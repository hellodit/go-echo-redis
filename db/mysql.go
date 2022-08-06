package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func ConnectGorm() *gorm.DB {

	connection := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}
