package filesystem

import (
	"os"
	"path/filepath"
	"server/api/dj/database/model"
	"server/api/dj/log"
	"strings"
)

func GetFiles(path string) ([]model.Song, error) {
	files := []model.Song{}

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Log(err)
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subdirPath := filepath.Join(path, entry.Name())
			subdirFiles, err := GetFiles(subdirPath)
			if err != nil {
				log.Log(err)
				return nil, err
			}
			files = append(files, subdirFiles...)
		}

		if !strings.HasSuffix(entry.Name(), ".mp3") {
			continue
		}

		p := filepath.Join(path, entry.Name())
		stats, err := os.Stat(p)
		if err != nil {
			log.Log(err)
			return nil, err
		}

		path := p
		datetime := stats.ModTime()
		file := model.Song{Path: path, Datetime: datetime}
		files = append(files, file)
	}

	return files, nil
}
