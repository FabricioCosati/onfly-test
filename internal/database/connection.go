package db

import (
	"fmt"
	"time"

	"github.com/FabricioCosati/onfly-test/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase(config config.Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		config.DbUser, config.DbPass, config.DbHost, config.DbPort, config.DbName, config.TimeZone,
	)

	for i := 0; i < 5; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			time.Sleep(8 * time.Second)
			continue
		}
		break
	}
	return db, err
}
