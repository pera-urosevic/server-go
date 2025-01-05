package database

import (
	"server/api/cleaner/types"
)

func GetPath(path string) ([]types.RecordCleaner, error) {
	db, err := Database()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM [cleaner] WHERE ? like [path] || '%'", path)
	if err != nil {
		return nil, err
	}

	records := []types.RecordCleaner{}
	for rows.Next() {
		record := types.RecordCleaner{}
		err := rows.Scan(&record.Path, &record.Name)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func AddPath(record types.RecordCleaner) error {
	db, err := Database()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO [cleaner] ([path], [name]) VALUES (?, ?)", record.Path, record.Name)
	if err != nil {
		return err
	}

	return nil
}

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
