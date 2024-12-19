package playlist

import (
	"server/api/dj/database"
	"strings"
)

func Playlist(q string) string {
	db, err := database.Database()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sql := database.GetQuery(db, q)
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	lines := []string{}
	lines = append(lines, "#EXTM3U")
	lines = append(lines, "")
	count := 0
	for rows.Next() {
		var path string
		var meta string
		var datetime string
		err := rows.Scan(&path, &meta, &datetime)
		if err != nil {
			panic(err)
		}
		count++
		lines = append(lines, path)
		lines = append(lines, "")
	}

	playlist := strings.Join(lines, "\n")
	return playlist
}
