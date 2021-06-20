package agent

import (
	"encoding/json"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/galenguyer/retina/client"
	"github.com/galenguyer/retina/core"
)

var (
	lock sync.Mutex
)

func Start() {
	services := []string{"https://example.com", "http://neverssl.com"}
	for _, service := range services {
		go monitor(service)
		time.Sleep(1 * time.Second)
	}
}

func monitor(service string) {
	for {
		lock.Lock()

		result := &core.Result{}
		if strings.HasPrefix(service, "https://") {
			result = client.PerformHTTPSCheck(service)
		} else if strings.HasPrefix(service, "http://") {
			result = client.PerformHTTPCheck(service)
		} else {
			log.Print("[ERROR][agent] invalid address", service)
		}

		js, _ := json.Marshal(result)
		log.Println(string(js))

		lock.Unlock()

		time.Sleep(5 * time.Second)
	}
}
