package database

import (
	"server/api/blog/database/model"
	"server/api/blog/log"
)

func GetPosts(filter string) ([]model.Post, error) {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return nil, err
	}

	posts := []model.Post{}
	f := "%" + filter + "%"
	db.Where("url LIKE ? OR title LIKE ? OR description LIKE ?", f, f, f).Order("timestamp DESC").Find(&posts)

	return posts, nil
}
