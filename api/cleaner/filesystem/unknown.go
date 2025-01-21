package filesystem

import (
	"os"
	"server/api/cleaner/log"
	"server/api/cleaner/types"
	"strings"
)

func isKnown(known []types.RecordCleaner, path string, name string) bool {
	for _, record := range known {
		if record.Path == path && record.Name == name {
			return true
		}
	}
	return false
}

func Unknown(known []types.RecordCleaner, scanPath string) ([]types.Node, error) {
	unknown := []types.Node{}
	root := os.Getenv("CLEANER_ROOT")
	base := strings.Replace(scanPath, root, "", 1)
	if base == "" {
		base = "/"
	}

	items, err := os.ReadDir(scanPath)
	if err != nil {
		if !strings.HasSuffix(err.Error(), ": Access is denied.") {
			log.Log(err)
			return nil, err
		}
	}

	for _, item := range items {
		name := item.Name()

		if isKnown(known, base, name) {
			continue
		}

		entry := types.Entry{}

		stats, err := os.Stat(scanPath + "/" + item.Name())
		if err != nil {
			log.Log(err)
			return nil, err
		}

		entry.Name = item.Name()

		children := []types.Node{}
		if stats.IsDir() {
			entry.Size = -1
			children, err = Unknown(known, scanPath+"/"+name)
		} else {
			entry.Size = stats.Size()
		}

		entry.Modified = stats.ModTime().Format("2006-01-02 15:04:05")

		if err != nil {
			log.Log(err)
			return nil, err
		}

		node := types.Node{Entry: entry, Children: children}
		unknown = append(unknown, node)
	}

	return unknown, nil
}
