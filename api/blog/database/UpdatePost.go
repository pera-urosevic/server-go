package database

import (
	"server/api/blog/log"
	"server/api/blog/types"
)

func UpdatePost(post types.Post) error {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE [blog] SET [timestamp]=?, [title]=?, [category]=?, [template]=?, [description]=?, [image]=?, [url]=? WHERE [id] = ?", post.Timestamp, post.Title, post.Category, post.Template, post.Description, post.Image, post.URL, post.ID)
	if err != nil {
		log.Log(err)
		return err
	}

	return nil
}
