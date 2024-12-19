package sync

import (
	"server/api/gallery/database"
	"server/api/gallery/log"
	"server/api/gallery/types"
)

func Sync() ([]types.Photo, error) {
	log.Log("SYNC", "started")
	records, err := database.GetPhotos("")
	if err != nil {
		return nil, err
	}
	files, err := scanFiles()
	if err != nil {
		return nil, err
	}
	records, err = checkOld(records, files)
	if err != nil {
		return nil, err
	}
	records, err = checkNew(records, files)
	if err != nil {
		return nil, err
	}
	log.Log("SYNC", "done")
	return records, nil
}
