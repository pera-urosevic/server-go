package flick

import (
	"errors"
	"server/api/gallery/types"
)

func Upload(photo types.Photo) ([]string, error) {
	messagesInit()
	messageLog("Flick")

	if photo.Path == "" {
		return nil, errors.New("path is required")
	}

	if photo.Title == "" {
		return nil, errors.New("title is required")
	}

	if photo.Keywords == "" {
		return nil, errors.New("keywords are required")
	}

	err := configLoad()
	if err != nil {
		return nil, err
	}

	client, err := auth()
	if err != nil {
		return nil, err
	}

	params := UploadParams{
		Path:        photo.Path,
		Title:       photo.Title,
		Description: photo.Description,
		Keywords:    photo.Keywords,
	}

	err = flickrUpload(client, params)
	if err != nil {
		return nil, err
	}

	messageLog("Done")

	return getMessages(), nil
}
