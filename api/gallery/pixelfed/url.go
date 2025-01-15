package pixelfed

import (
	"errors"
	"fmt"
	"os"
)

func getURL() (string, error) {
	server := os.Getenv("GALLERY_PIXELFED_SERVER")
	if server == "" {
		return "", errors.New("GALLERY_PIXELFED_SERVER not set")
	}

	url := fmt.Sprintf("%s/api/v2/media", server)
	return url, nil
}
