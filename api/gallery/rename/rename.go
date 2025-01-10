package rename

import (
	"os"
	"path/filepath"
	"server/api/gallery/types"
)

func Rename(photoOriginal types.Photo, photo types.Photo) (string, error) {
	oldPath := photoOriginal.Path
	newPath := filepath.Dir(photoOriginal.Path) + "\\" + photo.Datetime + " - " + photo.Title + "." + photo.Type

	if oldPath != newPath {
		err := os.Rename(oldPath, newPath)
		if err != nil {
			return photoOriginal.Path, err
		}

		return newPath, nil
	}

	return photoOriginal.Path, nil
}
