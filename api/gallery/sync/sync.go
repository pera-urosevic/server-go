package sync

import (
	"server/api/gallery/database"
	"server/api/gallery/database/model"
	"server/api/gallery/log"
)

func Sync() ([]model.Photo, error) {
	log.Log("[SYNC]", "started")
	records, err := database.GetPhotos("", "desc")

	if err != nil {
		log.Log(err)
		return nil, err
	}

	files, err := scanFiles()
	if err != nil {
		log.Log(err)
		return nil, err
	}

	records, err = checkOld(records, files)
	if err != nil {
		log.Log(err)
		return nil, err
	}

	records, err = checkNew(records, files)
	if err != nil {
		log.Log(err)
		return nil, err
	}

	log.Log("[SYNC]", "done")
	return records, nil
}
