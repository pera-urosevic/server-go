package pixelfed

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"server/api/gallery/types"
)

func getChecksum(photo types.Photo) string {
	f, err := os.Open(photo.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	checksum := fmt.Sprintf("%x", h.Sum(nil))
	return checksum
}
