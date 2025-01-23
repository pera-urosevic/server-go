package database

import (
	"server/api/gallery/database/model"
	"server/api/gallery/log"
)

func GetPhotos(filter string, sort string, pixelfedUpload bool) ([]model.Photo, error) {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return nil, err
	}

	photos := []model.Photo{}

	where := "(album LIKE ? OR datetime LIKE ? OR title LIKE ? OR description LIKE ? OR keywords LIKE ?)"
	if pixelfedUpload {
		where += "AND (pixelfed = '' OR pixelfed IS NULL) AND (online = 1)"
	}

	f := "%" + filter + "%"

	res := db.Where(where, f, f, f, f, f).Order("datetime " + sort).Find(&photos)
	if res.Error != nil {
		log.Log(res.Error)
		return nil, res.Error
	}

	return photos, nil
}
