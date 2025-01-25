package check

import (
	"net/http"
	"os"
	"server/micron/log"
	"time"
)

func Net() bool {
	url := os.Getenv("MICRON_CHECK_NET_URL")
	if url == "" {
		log.Log("env error, env MICRON_CHECK_NET_URL not found")
		return true
	}

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	_, err := client.Head(url)
	return err == nil
}
