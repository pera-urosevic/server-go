package database

import (
	"server/api/gallery/database/model"
	"server/api/gallery/log"
)

func GetPhotos(filter string, sort string) ([]model.Photo, error) {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return nil, err
	}

	f := "%" + filter + "%"
	photos := []model.Photo{}
	res := db.Where("album LIKE ? OR datetime LIKE ? OR title LIKE ? OR description LIKE ? OR keywords LIKE ?", f, f, f, f, f).Order("datetime " + sort).Find(&photos)
	if res.Error != nil {
		log.Log(res.Error)
		return nil, res.Error
	}

	return photos, nil
}
