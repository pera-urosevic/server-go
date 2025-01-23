package playlist

import (
	"server/api/dj/database"
	"server/api/dj/database/model"
	"server/api/dj/log"
	"strings"
)

func Playlist(q string) (string, error) {
	db, err := database.Database()
	if err != nil {
		log.Log(err)
		return "", err
	}

	query, err := database.GetQuery(db, q)
	if err != nil {
		log.Log(err)
		return "", err
	}

	songs := []model.Song{}
	res := db.Raw(query.Query).Scan(&songs)
	if res.Error != nil {
		log.Log(res.Error)
		return "", res.Error
	}

	lines := []string{}
	lines = append(lines, "#EXTM3U")
	lines = append(lines, "")
	count := 0
	for _, song := range songs {
		count++
		lines = append(lines, song.Path)
		lines = append(lines, "")
	}

	playlist := strings.Join(lines, "\n")
	return playlist, nil
}
