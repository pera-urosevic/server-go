package database

import (
	"server/api/gallery/database/model"
	"server/api/gallery/log"
)

func AddPhoto(photo model.Photo) (int64, error) {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return 0, err
	}

	res := db.Create(&photo)
	if res.Error != nil {
		log.Log(res.Error)
		return 0, res.Error
	}

	return photo.ID, nil
}
