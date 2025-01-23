package database

import (
	"server/api/gallery/database/model"
	"server/api/gallery/log"
)

func UpdatePhoto(photo model.Photo) error {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return err
	}

	res := db.Save(&photo)
	if res.Error != nil {
		log.Log(res.Error)
		return res.Error
	}

	return nil
}
