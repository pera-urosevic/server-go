package database

import (
	"server/api/cleaner/types"
)

func RemovePath(record types.RecordCleaner) error {
	db, err := Database()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM [cleaner] WHERE [path] = ? AND [name] = ?", record.Path, record.Name)
	if err != nil {
		return err
	}

	return nil
}
