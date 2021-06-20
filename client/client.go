package client

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func CheckWebsite(address string) {
	start := time.Now()
	resp, err := http.Get(address)
	if err != nil {
		log.Fatal(err)
	}
	duration := time.Since(start)

	if strings.HasPrefix(address, "http://") {
		log.Println("site:", address, "time:", duration, "status:", resp.StatusCode, "cert: none")

	} else {
		log.Println("site:", address, "time:", duration, "status:", resp.StatusCode, "cert:", time.Until(resp.TLS.PeerCertificates[0].NotAfter))

	}
}
