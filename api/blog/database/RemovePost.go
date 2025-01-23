package database

import (
	"server/api/blog/database/model"
	"server/api/blog/log"
)

func RemovePost(postID int64) error {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return err
	}

	db.Delete(&model.Post{}, 10)

	return nil
}
