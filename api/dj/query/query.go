package query

import (
	"encoding/json"
	"server/api/dj/database"
	"server/api/dj/log"
	"server/api/dj/types"
)

func Query(q string) ([]types.QueryResult, error) {
	db, err := database.Database()
	if err != nil {
		log.Log(err)
		return nil, err
	}

	query, err := database.GetQuery(db, q)
	if err != nil {
		log.Log(err)
		return nil, err
	}

	resultsRaw := []types.QueryResultRaw{}
	res := db.Raw(query.Query).Scan(&resultsRaw)
	if res.Error != nil {
		log.Log(res.Error)
		return nil, res.Error
	}

	results := []types.QueryResult{}
	for _, resultRaw := range resultsRaw {
		result := types.QueryResult{
			Query:    resultRaw.Query,
			Path:     resultRaw.Path,
			Datetime: resultRaw.Datetime,
		}
		json.Unmarshal([]byte(resultRaw.Meta), &result.Meta)
		results = append(results, result)
	}

	return results, nil
}
