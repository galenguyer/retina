package agent

import (
	"sync"
	"time"

	"github.com/galenguyer/retina/client"
)

var (
	lock sync.Mutex
)

func Start() {
	services := []string{"https://galenguyer.com", "https://vault.galenguyer.com"}
	for _, service := range services {
		time.Sleep(100 * time.Millisecond)
		go monitor(service)
	}
}

func monitor(service string) {
	for {
		lock.Lock()
		client.CheckWebsite(service)
		lock.Unlock()
		time.Sleep(5 * time.Second)
	}
}
