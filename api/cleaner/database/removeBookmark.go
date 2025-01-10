package database

func RemoveBookmark(bookmark string) error {
	db, err := Database()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM [bookmarks] WHERE [path] = ?", bookmark)
	if err != nil {
		return err
	}

	return nil
}
