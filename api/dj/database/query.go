package database

import "database/sql"

func GetQuery(db *sql.DB, name string) string {
	var query string
	row := db.QueryRow("SELECT query FROM queries WHERE name = ?", name)
	err := row.Scan(&query)
	if err != nil {
		return name
	}
	return query
}
