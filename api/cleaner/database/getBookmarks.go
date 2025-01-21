package database

import "server/api/cleaner/log"

func GetBookmarks() ([]string, error) {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM [bookmarks] order by [path]")
	if err != nil {
		log.Log(err)
		return nil, err
	}

	bookmarks := []string{}
	for rows.Next() {
		bookmark := ""
		err := rows.Scan(&bookmark)
		if err != nil {
			log.Log(err)
			return nil, err
		}
		bookmarks = append(bookmarks, bookmark)
	}

	return bookmarks, nil
}
