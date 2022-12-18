package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func migrate() {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		return
	}
}

func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database

	migrate()
}
