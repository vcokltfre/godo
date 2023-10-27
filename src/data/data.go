package data

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() error {
	dbLoc := os.Getenv("GODO_DB")
	if dbLoc == "" {
		dbLoc = "godo.db"
	}

	var err error
	db, err = gorm.Open(sqlite.Open(dbLoc), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.AutoMigrate(&Task{})
}
