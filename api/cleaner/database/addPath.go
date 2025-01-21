package database

import (
	"server/api/cleaner/log"
	"server/api/cleaner/types"
)

func AddPath(record types.RecordCleaner) error {
	db, err := Database()
	if err != nil {
		log.Log(err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO [cleaner] ([path], [name]) VALUES (?, ?)", record.Path, record.Name)
	if err != nil {
		log.Log(err)
		return err
	}

	return nil
}
