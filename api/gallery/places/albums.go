package places

import (
	"encoding/json"
	"os"
	"server/api/gallery/types"
	"server/env"
)

func albums() []types.Album {
	env.Load()
	data := os.Getenv("GALLERY_ALBUMS")
	if data == "" {
		panic("Missing GALLERY_ALBUMS")
	}
	albums := []types.Album{}
	json.Unmarshal([]byte(data), &albums)
	return albums
}

var Albums = albums()
