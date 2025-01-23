package query

import (
	"server/api/dj/database"
	"server/api/dj/database/model"
	"server/api/dj/log"
)

func Queries() ([]model.Query, error) {
	db, err := database.Database()
	if err != nil {
		log.Log(err)
		return nil, err
	}

	queries := []model.Query{}
	res := db.Find(&queries)
	if res.Error != nil {
		log.Log(res.Error)
		return nil, res.Error
	}

	return queries, nil
}
