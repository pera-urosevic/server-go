package database

func GetBookmarks() ([]string, error) {
	db, err := Database()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM [bookmarks] order by [path]")
	if err != nil {
		return nil, err
	}

	bookmarks := []string{}
	for rows.Next() {
		bookmark := ""
		err := rows.Scan(&bookmark)
		if err != nil {
			return nil, err
		}
		bookmarks = append(bookmarks, bookmark)
	}

	return bookmarks, nil
}

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
