package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

func Database() (*gorm.DB, error) {
	dbPath := os.Getenv("CLEANER_DB_PATH")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	return db, err
}
