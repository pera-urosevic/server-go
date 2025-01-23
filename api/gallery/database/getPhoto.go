package database

import (
	"server/api/gallery/database/model"
	"server/api/gallery/log"
)

func GetPhoto(photoID int64) (model.Photo, error) {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return model.Photo{}, err
	}

	photo := model.Photo{}
	res := db.First(&photo, photoID)
	if res.Error != nil {
		log.Log(res.Error)
		return model.Photo{}, res.Error
	}

	return photo, nil
}
