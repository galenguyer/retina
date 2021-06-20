package client

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/galenguyer/retina/core"
)

func CheckWebsite(address string) *core.Result {
	start := time.Now()
	resp, err := http.Get(address)
	if err != nil {
		log.Fatal(err)
	}
	duration := time.Since(start)

	result := &core.Result{HTTPStatusCode: resp.StatusCode, Duration: duration, Timestamp: start, URL: address}
	if strings.HasPrefix(address, "https://") {
		// cert expiration check here
	}
	return result
}
