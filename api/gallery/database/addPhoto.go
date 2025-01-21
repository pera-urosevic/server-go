package database

import (
	"server/api/gallery/log"
	"server/api/gallery/types"
)

func AddPhoto(photo types.Photo) (int64, error) {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return 0, err
	}
	defer db.Close()

	res, err := db.Exec("INSERT INTO [gallery] ([path], [type], [modified], [online], [album], [datetime], [title], [description], [keywords], [copyright], [flickr], [pixelfed]) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", photo.Path, photo.Type, photo.Modified, photo.Online, photo.Album, photo.Datetime, photo.Title, photo.Description, photo.Keywords, photo.Copyright, photo.Flickr, photo.PixelFed)
	if err != nil {
		log.Log(err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Log(err)
		return 0, err
	}

	return id, nil
}
