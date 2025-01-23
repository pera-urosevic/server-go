package search

import (
	"encoding/json"
	"server/api/dj/database"
	"server/api/dj/database/model"
	"server/api/dj/log"
	"server/api/dj/types"
)

func Search(q string) ([]types.QueryResult, error) {
	db, err := database.Database()
	if err != nil {
		log.Log(err)
		return nil, err
	}

	songs := []model.Song{}
	res := db.Find(&songs, "path LIKE ? OR meta LIKE ?", "%"+q+"%", "%"+q+"%")
	if res.Error != nil {
		log.Log(res.Error)
		return nil, res.Error
	}

	results := []types.QueryResult{}
	for _, song := range songs {
		result := types.QueryResult{
			Query:    "Search",
			Path:     song.Path,
			Datetime: song.Datetime,
		}
		json.Unmarshal([]byte(song.Meta), &result.Meta)
		results = append(results, result)
	}

	return results, nil
}
