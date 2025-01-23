package rename

import (
	"os"
	"path/filepath"
	"server/api/gallery/database/model"
	"server/api/gallery/log"
)

func Rename(photoOriginal model.Photo, photo model.Photo) (string, error) {
	oldPath := photoOriginal.Path
	newPath := filepath.Dir(photoOriginal.Path) + "\\" + photo.Datetime + " - " + photo.Title + "." + photo.Type

	if oldPath != newPath {
		err := os.Rename(oldPath, newPath)
		if err != nil {
			log.Log(err)
			return photoOriginal.Path, err
		}

		return newPath, nil
	}

	return photoOriginal.Path, nil
}
