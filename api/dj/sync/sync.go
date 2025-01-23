package sync

import (
	"os"
	"server/api/dj/database"
	"server/api/dj/filesystem"
	"server/api/dj/log"
)

func Sync() (map[string][]string, error) {
	db, err := database.Database()
	if err != nil {
		log.Log(err)
		return nil, err
	}

	musicPath := os.Getenv("DJ_MUSIC_PATH")
	files, err := filesystem.GetFiles(musicPath)
	if err != nil {
		log.Log(err)
		return nil, err
	}

	records, err := getRecords(db)
	if err != nil {
		log.Log(err)
		return nil, err
	}

	records, removed, err := removeRecords(db, records, files)
	if err != nil {
		log.Log(err)
		return nil, err
	}

	records, added, err := addRecords(db, records, files)
	if err != nil {
		log.Log(err)
		return nil, err
	}

	if len(records) == 0 {
		log.Log("Synced to empty")
	}

	res := map[string][]string{
		"removed": removed,
		"added":   added,
	}
	return res, nil
}
