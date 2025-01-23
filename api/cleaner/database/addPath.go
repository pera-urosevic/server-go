package database

import (
	"server/api/cleaner/database/model"
	"server/api/cleaner/log"
)

func AddPath(record model.Cleaner) error {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return err
	}

	res := db.Create(&record)
	if res.Error != nil {
		log.Log(res.Error)
		return res.Error
	}

	return nil
}
