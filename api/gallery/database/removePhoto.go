package database

import "server/api/gallery/log"

func RemovePhoto(postID int64) error {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return err
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM [gallery] WHERE [id] = ?", postID)
	if err != nil {
		log.Log(err)
		return err
	}
	return nil
}
