package sync

import (
	"os"
	"path/filepath"
	"regexp"
	"server/api/gallery/places"
	"server/api/gallery/types"
)

func isPhoto(filename string) bool {
	match, err := regexp.MatchString(".+\\.(jpg|jpeg|png|gif)", filename)
	if err != nil {
		panic(err)
	}
	return match
}

func scanFiles() ([]types.AlbumFile, error) {
	files := []types.AlbumFile{}
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	for _, album := range places.Albums {
		absolutePath := filepath.Join(home, album.Path)
		entry, err := os.ReadDir(absolutePath)
		if err != nil {
			return nil, err
		}
		for _, file := range entry {
			info, err := file.Info()
			if err != nil {
				return nil, err
			}
			if isPhoto(info.Name()) {
				files = append(files, types.AlbumFile{
					Path:     filepath.Join(absolutePath, info.Name()),
					Album:    album.Title,
					Modified: info.ModTime().UnixMilli(),
					Online:   album.Online,
				})
			}
		}
	}
	return files, nil
}
