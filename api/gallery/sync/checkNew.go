package sync

import (
	"fmt"
	"server/api/gallery/database"
	"server/api/gallery/database/model"
	"server/api/gallery/log"
	"server/api/gallery/places"
	"server/api/gallery/types"

	"github.com/disintegration/imaging"
)

func addRecord(file types.AlbumFile) (model.Photo, error) {
	record, err := readMeta(file)
	if err != nil {
		log.Log(err)
		return record, err
	}

	id, err := database.AddPhoto(record)
	if err != nil {
		log.Log(err)
		return record, err
	}

	record.ID = id
	return record, nil
}

func addThumbnails(record model.Photo) error {
	original, err := imaging.Open(record.Path)
	if err != nil {
		log.Log(err)
		return err
	}

	filename := fmt.Sprintf("%s - %s.%s", record.Album, record.Datetime, record.Type)
	image := imaging.Fit(original, 1600, 900, imaging.Lanczos)
	imagePath := fmt.Sprintf("%s\\%s", places.ImagesCache, filename)
	err = imaging.Save(image, imagePath, imaging.JPEGQuality(75))
	if err != nil {
		log.Log(err)
		return err
	}

	preview := imaging.Fit(original, 360, 360, imaging.Lanczos)
	previewPath := fmt.Sprintf("%s\\%s", places.PreviewsCache, filename)
	err = imaging.Save(preview, previewPath, imaging.JPEGQuality(75))
	if err != nil {
		log.Log(err)
		return err
	}

	thumbnail := imaging.Fill(original, 160, 140, imaging.Center, imaging.Lanczos)
	thumbnailPath := fmt.Sprintf("%s\\%s", places.ThumbnailsCache, filename)
	err = imaging.Save(thumbnail, thumbnailPath, imaging.JPEGQuality(75))
	if err != nil {
		log.Log(err)
		return err
	}

	return nil
}

func checkNew(records []model.Photo, files []types.AlbumFile) ([]model.Photo, error) {
	for _, file := range files {
		found := false

		for _, record := range records {
			if file.Path == record.Path && file.Modified == record.Modified {
				found = true
				break
			}
		}

		if !found {
			log.Log("[ADD]", file.Path)
			recordNew, err := addRecord(file)
			if err != nil {
				log.Log(err)
				return nil, err
			}

			err = addThumbnails(recordNew)
			if err != nil {
				log.Log(err)
				return nil, err
			}

			records = append(records, recordNew)
		}
	}

	return records, nil
}
