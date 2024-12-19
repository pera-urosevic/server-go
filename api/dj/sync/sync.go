package sync

import (
	"os"
	"server/api/dj/database"
	"server/api/dj/filesystem"
	"server/api/dj/log"
)

func Sync() map[string][]string {
	db, err := database.Database()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	musicPath := os.Getenv("DJ_MUSIC_PATH")
	files := filesystem.GetFiles(musicPath)
	records := getRecords(db)
	records, removed := removeRecords(db, records, files)
	records, added := addRecords(db, records, files)

	if len(records) == 0 {
		log.Log("Synced to empty")
	}

	res := map[string][]string{
		"removed": removed,
		"added":   added,
	}
	return res
}
