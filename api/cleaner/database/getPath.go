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
