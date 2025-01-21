package database

import "server/api/cleaner/log"

func RemoveBookmark(bookmark string) error {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM [bookmarks] WHERE [path] = ?", bookmark)
	if err != nil {
		log.Log(err)
		return err
	}

	return nil
}
