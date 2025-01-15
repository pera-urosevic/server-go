package pixelfed

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"server/api/gallery/types"
	"time"
)

func Upload(photo types.Photo) (string, error) {
	url, err := getURL()
	if err != nil {
		return "", err
	}

	body, err := getBody(photo)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return "", err
	}

	bearerToken, err := getBearerToken()
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", bearerToken)

	checksum := getChecksum(photo)
	req.Header.Set("Idempotency-Key", checksum)

	client := http.Client{
		Timeout: 2 * time.Minute,
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", errors.New("Upload status: " + res.Status)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// TODO: add to collection
	// curl --silent --header "Authorization: Bearer $ACCESSTOKEN" -X POST --form "media_ids[]=$media_id" https://$SERVER/api/v1/statuses | jq '.url'
	return "", nil
}
