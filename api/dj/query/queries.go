package query

import (
	"server/api/dj/database"
	"server/api/dj/types"
)

func Queries() map[string]string {
	db, err := database.Database()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, query FROM queries")
	if err != nil {
		panic(err)
	}

	queries := make(map[string]string)
	for rows.Next() {
		result := types.RecordQuery{}
		err := rows.Scan(&result.Name, &result.Query)
		if err != nil {
			panic(err)
		}
		queries[result.Name] = result.Query
	}

	return queries
}
