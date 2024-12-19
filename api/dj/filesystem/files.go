package filesystem

import (
	"os"
	"path/filepath"
	"server/api/dj/types"
	"strings"
)

func GetFiles(path string) []types.RecordSong {
	files := []types.RecordSong{}

	entries, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subdirPath := filepath.Join(path, entry.Name())
			subdirFiles := GetFiles(subdirPath)
			files = append(files, subdirFiles...)
		}

		if !strings.HasSuffix(entry.Name(), ".mp3") {
			continue
		}

		p := filepath.Join(path, entry.Name())
		stats, err := os.Stat(p)
		if err != nil {
			panic(err)
		}

		path := p
		datetime := stats.ModTime()
		file := types.RecordSong{Path: path, Datetime: datetime}
		files = append(files, file)
	}

	return files
}
