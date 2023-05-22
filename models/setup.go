package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=addie dbname=go_blog port=5432 sslmode=disable timezone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) //change DB provider if necessary

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Post{})

	DB = database
}