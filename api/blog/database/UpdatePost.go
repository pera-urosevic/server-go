package database

import (
	"server/api/blog/database/model"
	"server/api/blog/log"
)

func UpdatePost(post model.Post) error {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return err
	}

	res := db.Save(&post)
	if res.Error != nil {
		log.Log(res.Error)
		return res.Error
	}

	return nil
}
