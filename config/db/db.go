package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateDatabase() {
	db, err := gorm.Open(sqlite.Open("./config/db/database.db"), &gorm.Config{})

	if err != nil {
		panic("failed to cconnect database")
	}

	DB = db
}