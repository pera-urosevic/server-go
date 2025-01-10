package database

import (
	"os"

	"database/sql"

	_ "modernc.org/sqlite"
)

func Database() (*sql.DB, error) {
	dbPath := os.Getenv("BLOG_DB_PATH")
	db, err := sql.Open("sqlite", dbPath)
	return db, err
}
