package database

import (
	"server/api/cleaner/database/model"
	"server/api/cleaner/log"
)

func GetPaths() ([]model.Cleaner, error) {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return nil, err
	}

	records := []model.Cleaner{}
	res := db.Find(&records)
	if res.Error != nil {
		log.Log(res.Error)
		return nil, res.Error
	}

	return records, nil
}
