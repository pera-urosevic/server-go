package search

import (
	"encoding/json"
	"fmt"
	"server/api/dj/database"
	"server/api/dj/types"
)

func Search(q string) []types.RecordQueryResult {
	db, err := database.Database()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sql := fmt.Sprintf("SELECT 'Search' as query, * FROM songs WHERE path LIKE '%%%s%%' OR meta LIKE '%%%s%%'", q, q)

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	results := []types.RecordQueryResult{}
	for rows.Next() {
		result := types.RecordQueryResult{}
		var metaJson string
		err := rows.Scan(&result.Query, &result.Path, &metaJson, &result.Datetime)
		json.Unmarshal([]byte(metaJson), &result.Meta)
		if err != nil {
			panic(err)
		}
		results = append(results, result)
	}

	return results
}
