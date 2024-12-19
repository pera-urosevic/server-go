package sync

import (
	"database/sql"
	"encoding/json"
	"server/api/dj/filesystem"
	"server/api/dj/types"
)

func getRecords(db *sql.DB) []types.RecordSong {
	rows, err := db.Query("SELECT path, datetime FROM songs")
	if err != nil {
		panic(err)
	}

	records := []types.RecordSong{}
	for rows.Next() {
		record := types.RecordSong{}
		err := rows.Scan(&record.Path, &record.Datetime)
		if err != nil {
			panic(err)
		}
		records = append(records, record)
	}

	return records
}

func removeRecords(db *sql.DB, records []types.RecordSong, files []types.RecordSong) ([]types.RecordSong, []string) {
	removed := []string{}
	foundRecords := []types.RecordSong{}
	for _, record := range records {
		found := false
		for _, file := range files {
			if record.Path == file.Path && record.Datetime == file.Datetime {
				found = true
				break
			}
		}
		if found {
			foundRecords = append(foundRecords, record)
		} else {
			removed = append(removed, record.Path)
			_, err := db.Exec("DELETE FROM [songs] WHERE [path] = ?", record.Path)
			if err != nil {
				panic(err)
			}
		}
	}
	return foundRecords, removed
}

func addRecords(db *sql.DB, records []types.RecordSong, files []types.RecordSong) ([]types.RecordSong, []string) {
	added := []string{}
	for _, file := range files {
		found := false
		for _, record := range records {
			if record.Path == file.Path {
				found = true
				break
			}
		}
		if !found {
			added = append(added, file.Path)
			file.Meta = filesystem.ReadMeta(file.Path)
			jsonMeta, err := json.Marshal(file.Meta)
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("INSERT INTO [songs] ([path], [meta], [datetime]) VALUES (?, ?, ?)", file.Path, string(jsonMeta), file.Datetime)
			if err != nil {
				panic(err)
			}
			records = append(records, file)
		}
	}
	return records, added
}
