package database

import (
	"server/api/cleaner/database/model"
	"server/api/cleaner/log"
)

func AddBookmark(bookmark string) error {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return err
	}

	res := db.Create(&model.Bookmark{Path: bookmark})
	if res.Error != nil {
		log.Log(res.Error)
		return res.Error
	}

	return nil
}
