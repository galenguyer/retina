package client

import (
	"log"
	"net/http"
	"time"

	"github.com/galenguyer/retina/core"
)

func PerformHTTPCheck(address string) *core.Result {
	start := time.Now()
	resp, err := http.Get(address)
	if err != nil {
		log.Fatal(err)
	}
	duration := time.Since(start)

	result := &core.Result{HTTPStatusCode: resp.StatusCode, Duration: duration, Timestamp: start, URL: address}
	return result
}

func PerformHTTPSCheck(address string) *core.Result {
	start := time.Now()
	resp, err := http.Get(address)
	if err != nil {
		log.Fatal(err)
	}
	duration := time.Since(start)

	result := &core.Result{
		HTTPStatusCode:    resp.StatusCode,
		Duration:          duration,
		Timestamp:         start,
		URL:               address,
		CertificateExpiry: time.Until(resp.TLS.PeerCertificates[0].NotAfter)}
	return result
}
