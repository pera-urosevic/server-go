package database

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
