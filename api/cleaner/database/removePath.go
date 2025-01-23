package database

import (
	"server/api/cleaner/database/model"
	"server/api/cleaner/log"
)

func RemovePath(record model.Cleaner) error {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return err
	}

	res := db.Where("path = ? AND name = ?", record.Path, record.Name).Delete(&model.Cleaner{})
	if res.Error != nil {
		log.Log(res.Error)
		return res.Error
	}

	return nil
}
