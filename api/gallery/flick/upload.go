package flick

import (
	"server/system"
	"strings"

	"gopkg.in/masci/flickr.v3"
)

func flickrUpload(client *flickr.FlickrClient, uploadParams UploadParams) error {
	system.Log("Uploading: " + uploadParams.Path)

	client.Init()
	params := flickr.NewUploadParams()
	params.Title = uploadParams.Title
	params.Description = uploadParams.Description
	params.Tags = strings.Split(uploadParams.Keywords, " | ")
	params.IsPublic = true
	params.IsFriend = false
	params.IsFamily = false
	res, err := flickr.UploadFile(client, uploadParams.Path, params)
	if err != nil {
		return err
	} else {
		system.Log("Uploaded: " + res.ID)
	}

	return nil
}
