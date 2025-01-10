package database

func RemovePhoto(postID int64) error {
	db, err := Database()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM [gallery] WHERE [id] = ?", postID)
	if err != nil {
		return err
	}
	return nil
}
