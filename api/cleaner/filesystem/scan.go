package filesystem

import (
	"os"
	"server/api/cleaner/log"
	"server/api/cleaner/types"
)

func Scan(path string) ([]types.Entry, error) {
	entries := []types.Entry{}

	root := os.Getenv("CLEANER_ROOT")

	scanPath := root + path

	items, err := os.ReadDir(scanPath)
	if err != nil {
		log.Log(err)
		return nil, err
	}

	for _, item := range items {
		entry := types.Entry{}

		stats, err := os.Stat(scanPath + "/" + item.Name())
		if err != nil {
			log.Log(err)
			return nil, err
		}

		entry.Name = item.Name()

		if stats.IsDir() {
			entry.Size = -1
		} else {
			entry.Size = stats.Size()
		}

		entry.Modified = stats.ModTime().Format("2006-01-02 15:04:05")

		entries = append(entries, entry)
	}

	return entries, nil
}
