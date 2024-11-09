package db

import (
	"fmt"

	"github.com/FabricioCosati/onfly-test/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		config.DbUser, config.DbPass, config.DbHost, config.DbPort, config.DbName, config.TimeZone,
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
