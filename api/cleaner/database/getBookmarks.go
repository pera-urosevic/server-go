package database

import (
	"server/api/cleaner/database/model"
	"server/api/cleaner/log"
)

func GetBookmarks() ([]string, error) {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return nil, err
	}

	bookmarks := []model.Bookmark{}
	res := db.Find(&bookmarks)
	if res.Error != nil {
		log.Log(res.Error)
		return nil, res.Error
	}

	bookmarkPaths := []string{}
	for _, bookmark := range bookmarks {
		bookmarkPaths = append(bookmarkPaths, bookmark.Path)
	}

	return bookmarkPaths, nil
}
