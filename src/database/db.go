package database

import (
	"go-ambassador/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open(mysql.Open("root:root@tcp(db:3306)/ambassador"), &gorm.Config{})

	if err != nil {
		panic("colud not connect with the database")
	}

}

func AutoMigrate() {
	DB.AutoMigrate(models.User{})
}
