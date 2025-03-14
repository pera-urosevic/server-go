package sync

import (
	"fmt"
	"os"
	"server/api/gallery/database"
	"server/api/gallery/database/model"
	"server/api/gallery/log"
	"server/api/gallery/places"
	"server/api/gallery/types"
)

func removeThumbnails(record model.Photo) error {
	filename := fmt.Sprintf("%s - %s.%s", record.Album, record.Datetime, record.Type)
	err := os.Remove(fmt.Sprintf("%s\\%s", places.ImagesCache, filename))
	if err != nil {
		log.Log(err)
		return err
	}

	err = os.Remove(fmt.Sprintf("%s\\%s", places.PreviewsCache, filename))
	if err != nil {
		log.Log(err)
		return err
	}

	err = os.Remove(fmt.Sprintf("%s\\%s", places.ThumbnailsCache, filename))
	if err != nil {
		log.Log(err)
		return err
	}

	return nil
}

func removeDB(photo model.Photo) error {
	err := database.RemovePhoto(photo.ID)
	if err != nil {
		log.Log(err)
	}
	return err
}

func checkOld(records []model.Photo, files []types.AlbumFile) ([]model.Photo, error) {
	checked := []model.Photo{}
	for _, record := range records {

		found := false
		for _, file := range files {
			if record.Path == file.Path && file.Modified == record.Modified {
				found = true
				checked = append(checked, record)
				break
			}
		}

		if !found {
			log.Log("[REMOVE]", record.Path)
			err := removeThumbnails(record)
			if err != nil {
				log.Log(err)
				return nil, err
			}

			err = removeDB(record)
			if err != nil {
				log.Log(err)
				return nil, err
			}
		}
	}

	return checked, nil
}
