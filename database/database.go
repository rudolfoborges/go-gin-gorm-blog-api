package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	DB, err = gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
