package pixelfed

import (
	"fmt"
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

	// --header "Idempotency-Key: $checksum"

	client := http.Client{
		Timeout: 2 * time.Minute,
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	fmt.Println("// DEBUG! res", res)

	// TODO! collection
	// curl --silent --header "Authorization: Bearer $ACCESSTOKEN" -X POST --form "media_ids[]=$media_id" https://$SERVER/api/v1/statuses | jq '.url'
	return "", nil
}
