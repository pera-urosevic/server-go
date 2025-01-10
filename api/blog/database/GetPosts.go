package database

import "server/api/blog/types"

func GetPosts(filter string) ([]types.Post, error) {
	db, err := Database()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	f := "%" + filter + "%"
	rows, err := db.Query("SELECT * FROM [blog] WHERE [url] LIKE ? OR [title] LIKE ? OR [description] LIKE ? ORDER BY [timestamp] DESC", f, f, f)
	if err != nil {
		return nil, err
	}

	posts := []types.Post{}
	for rows.Next() {
		post := types.Post{}
		err := rows.Scan(&post.ID, &post.Timestamp, &post.Title, &post.Category, &post.Template, &post.Description, &post.Image, &post.URL)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
