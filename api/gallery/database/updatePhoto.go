package database

import "server/api/gallery/types"

func UpdatePhoto(photo types.Photo) error {
	db, err := Database()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("UPDATE [gallery] SET [path]=?, [type]=?, [modified]=?, [online]=?, [album]=?, [datetime]=?, [title]=?, [description]=?, [keywords]=?, [copyright]=?, [flickr]=?, [pixelfed]=? WHERE [id] = ?", photo.Path, photo.Type, photo.Modified, photo.Online, photo.Album, photo.Datetime, photo.Title, photo.Description, photo.Keywords, photo.Copyright, photo.Flickr, photo.PixelFed, photo.ID)
	if err != nil {
		return err
	}
	return nil
}
