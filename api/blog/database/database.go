package database

import (
	"os"
	"server/api/blog/types"
	"time"

	"database/sql"

	_ "modernc.org/sqlite"
)

func Database() (*sql.DB, error) {
	dbPath := os.Getenv("BLOG_DB_PATH")
	db, err := sql.Open("sqlite", dbPath)
	return db, err
}

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

func GetPost(postID int64) (types.Post, error) {
	db, err := Database()
	if err != nil {
		return types.Post{}, err
	}
	defer db.Close()
	post := types.Post{}
	row := db.QueryRow("SELECT * FROM [blog] WHERE [id] = ?", postID)
	err = row.Scan(&post.ID, &post.Timestamp, &post.Title, &post.Category, &post.Template, &post.Description, &post.Image, &post.URL)
	if err != nil {
		return types.Post{}, err
	}
	return post, nil
}

func InsertPost(url string) (int64, error) {
	db, err := Database()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	timestamp := time.Now().UnixMilli()
	res, err := db.Exec("INSERT INTO [blog] ([timestamp], [title], [category], [template], [description], [image], [url]) VALUES (?, ?, ?, ?, ?, ?, ?)", timestamp, "", "", "", "", "", url)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func UpdatePost(post types.Post) error {
	db, err := Database()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("UPDATE [blog] SET [timestamp]=?, [title]=?, [category]=?, [template]=?, [description]=?, [image]=?, [url]=? WHERE [id] = ?", post.Timestamp, post.Title, post.Category, post.Template, post.Description, post.Image, post.URL, post.ID)
	if err != nil {
		return err
	}
	return nil
}

func RemovePost(postID int64) error {
	db, err := Database()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM [blog] WHERE [id] = ?", postID)
	if err != nil {
		return err
	}
	return nil
}
