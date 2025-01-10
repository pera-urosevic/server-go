package database

func AddBookmark(bookmark string) error {
	db, err := Database()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO [bookmarks] ([path]) VALUES (?)", bookmark)
	if err != nil {
		return err
	}

	return nil
}
