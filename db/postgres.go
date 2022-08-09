package db

import (
	"github.com/sirupsen/logrus"
	"go-echo-redis/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func ConnectGorm() *gorm.DB {

	connection := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	db.AutoMigrate(domain.Author{}, domain.Post{})
	if err != nil {
		logrus.Error(err)
	}
	return db
}
