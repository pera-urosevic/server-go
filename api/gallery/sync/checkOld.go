package sync

import (
	"fmt"
	"os"
	"server/api/gallery/database"
	"server/api/gallery/log"
	"server/api/gallery/places"
	"server/api/gallery/types"
)

func removeThumbnails(record types.Photo) error {
	filename := fmt.Sprintf("%s - %s.%s", record.Album, record.Datetime, record.Type)
	err := os.Remove(fmt.Sprintf("%s\\%s", places.ImagesCache, filename))
	if err != nil {
		return err
	}
	err = os.Remove(fmt.Sprintf("%s\\%s", places.PreviewsCache, filename))
	if err != nil {
		return err
	}
	err = os.Remove(fmt.Sprintf("%s\\%s", places.ThumbnailsCache, filename))
	if err != nil {
		return err
	}
	return nil
}

func removeDB(photo types.Photo) error {
	err := database.RemovePhoto(photo.ID)
	return err
}

func checkOld(records []types.Photo, files []types.AlbumFile) ([]types.Photo, error) {
	checked := []types.Photo{}
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
			log.Log("REMOVE", record.Path)
			err := removeThumbnails(record)
			if err != nil {
				return nil, err
			}
			err = removeDB(record)
			if err != nil {
				return nil, err
			}
		}
	}
	return checked, nil
}
