package client

import (
	"log"
	"net/http"
	"time"
)

func CheckWebsite(address string) {
	start := time.Now()
	resp, err := http.Get(address)
	if err != nil {
		log.Fatal(err)
	}
	duration := time.Since(start)

	log.Println("site:", address, "time:", duration, "status:", resp.StatusCode)
}
