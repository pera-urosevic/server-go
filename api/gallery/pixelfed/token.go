package pixelfed

import (
	"errors"
	"os"
)

func getBearerToken() (string, error) {
	token := os.Getenv("GALLERY_PIXELFED_TOKEN")
	if token == "" {
		return "", errors.New("GALLERY_PIXELFED_TOKEN not set")
	}

	return "Bearer " + token, nil
}
