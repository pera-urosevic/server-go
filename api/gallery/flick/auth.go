package flick

import (
	"fmt"

	"gopkg.in/masci/flickr.v3"
)

func auth() (*flickr.FlickrClient, error) {

	client := flickr.NewFlickrClient(config.Api.Key, config.Api.Secret)

	if !hasFlickrAccess() {
		requestTok, _ := flickr.GetRequestToken(client)
		url, _ := flickr.GetAuthorizeUrl(client, requestTok)
		fmt.Println("Please visit the following URL and paste the code:\n", url)

		var confirm string
		fmt.Scanln(&confirm)

		accessTok, err := flickr.GetAccessToken(client, requestTok, confirm)
		if err != nil {
			return nil, err
		}

		client.OAuthToken = accessTok.OAuthToken
		client.OAuthTokenSecret = accessTok.OAuthTokenSecret
		config.Access.Token = accessTok.OAuthToken
		config.Access.Secret = accessTok.OAuthTokenSecret

		configSave()

	} else {
		client.OAuthToken = config.Access.Token
		client.OAuthTokenSecret = config.Access.Secret
	}

	return client, nil
}
