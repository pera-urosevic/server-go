package database

import (
	"server/api/gallery/log"
	"server/api/gallery/types"
)

func GetPhoto(photoID int64) (types.Photo, error) {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return types.Photo{}, err
	}
	defer db.Close()
	photo := types.Photo{}
	row := db.QueryRow("SELECT * FROM [gallery] WHERE [id] = ?", photoID)
	err = row.Scan(&photo.ID, &photo.Path, &photo.Type, &photo.Modified, &photo.Online, &photo.Album, &photo.Datetime, &photo.Title, &photo.Description, &photo.Keywords, &photo.Copyright, &photo.Flickr, &photo.PixelFed)
	if err != nil {
		log.Log(err)
		return types.Photo{}, err
	}
	return photo, nil
}
