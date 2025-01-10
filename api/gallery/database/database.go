package database

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

func Database() (*sql.DB, error) {
	dbPath := os.Getenv("GALLERY_DB_PATH")
	db, err := sql.Open("sqlite", dbPath)
	return db, err
}
