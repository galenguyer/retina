package agent

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/galenguyer/retina/client"
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
		result := client.CheckWebsite(service)
		js, _ := json.Marshal(result)
		log.Println(string(js))
		lock.Unlock()
		time.Sleep(5 * time.Second)
	}
}
