package pixelfed

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"server/api/gallery/types"
)

func getBody(photo types.Photo) (*bytes.Buffer, error) {
	file, err := os.Open(photo.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(photo.Path))
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	description := getDescription(photo)
	writer.WriteField("description", description)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}
