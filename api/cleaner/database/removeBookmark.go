package database

import (
	"server/api/cleaner/database/model"
	"server/api/cleaner/log"
)

func RemoveBookmark(path string) error {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return err
	}

	res := db.Where("path = ?", path).Delete(&model.Bookmark{})
	if res.Error != nil {
		log.Log(res.Error)
		return res.Error
	}

	return nil
}
