package configs

import (
	"fmt"
	"project/mock_api/models"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConnect(config *AppConfig) *gorm.DB {
	uri := ""

	uri = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name)

	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database: ", err)
	}

	DatabaseMigration(db)

	return db
}

func DatabaseMigration(db *gorm.DB) {
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Book{})
}
