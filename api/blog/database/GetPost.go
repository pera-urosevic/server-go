package database

import (
	"server/api/blog/log"
	"server/api/blog/types"
)

func GetPost(postID int64) (types.Post, error) {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return types.Post{}, err
	}
	defer db.Close()
	post := types.Post{}
	row := db.QueryRow("SELECT * FROM [blog] WHERE [id] = ?", postID)
	err = row.Scan(&post.ID, &post.Timestamp, &post.Title, &post.Category, &post.Template, &post.Description, &post.Image, &post.URL)
	if err != nil {
		log.Log(err)
		return types.Post{}, err
	}
	return post, nil
}
