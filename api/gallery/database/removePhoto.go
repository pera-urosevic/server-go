package database

import (
	"server/api/gallery/database/model"
	"server/api/gallery/log"
)

func RemovePhoto(postID int64) error {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return err
	}

	res := db.Delete(&model.Photo{}, postID)
	if res.Error != nil {
		log.Log(res.Error)
		return res.Error
	}

	return nil
}
