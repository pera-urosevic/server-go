package database

import (
	"server/api/blog/log"
	"time"
)

func InsertPost(url string) (int64, error) {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return 0, err
	}
	defer db.Close()

	timestamp := time.Now().UnixMilli()
	res, err := db.Exec("INSERT INTO [blog] ([timestamp], [title], [category], [template], [description], [image], [url]) VALUES (?, ?, ?, ?, ?, ?, ?)", timestamp, "", "", "", "", "", url)
	if err != nil {
		log.Log(err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Log(err)
		return 0, err
	}

	return id, nil
}
