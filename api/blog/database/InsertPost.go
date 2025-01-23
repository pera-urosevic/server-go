package database

import (
	"fmt"
	"server/api/blog/database/model"
	"server/api/blog/log"
	"time"
)

func InsertPost(url string) (int64, error) {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return 0, err
	}

	timestamp := time.Now().UnixMilli()
	post := model.Post{
		Timestamp:   fmt.Sprintf("%d", timestamp),
		Title:       "",
		Category:    "",
		Template:    "",
		Description: "",
		Image:       "",
		URL:         url,
	}

	res := db.Create(&post)
	if res.Error != nil {
		log.Log(res.Error)
		return 0, res.Error
	}

	return post.ID, nil
}
