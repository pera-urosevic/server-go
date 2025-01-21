package database

import "server/api/cleaner/log"

func AddBookmark(bookmark string) error {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO [bookmarks] ([path]) VALUES (?)", bookmark)
	if err != nil {
		log.Log(err)
		return err
	}

	return nil
}
