module go-echo-redis

// +heroku goVersion go1.16
go 1.16

require (
	github.com/go-redis/redis/v7 v7.4.0
	github.com/labstack/echo/v4 v4.2.2
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.7.1
	gorm.io/driver/postgres v1.3.5
	gorm.io/gorm v1.23.4
)
