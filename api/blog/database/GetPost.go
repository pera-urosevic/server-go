package database

import (
	"server/api/blog/database/model"
	"server/api/blog/log"
)

func GetPost(postID int64) (model.Post, error) {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return model.Post{}, err
	}

	post := model.Post{}
	db.First(&post, postID)

	return post, nil
}
