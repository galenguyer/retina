package agent

import (
	"encoding/json"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/galenguyer/retina/client"
	"github.com/galenguyer/retina/config"
	"github.com/galenguyer/retina/core"
)

var (
	lock sync.Mutex
)

func Start(config *config.Config) {
	for _, service := range config.Services {
		go monitor(service)
		time.Sleep(1 * time.Second)
	}
}

func monitor(service core.Service) {
	for {
		lock.Lock()

		result := &core.Result{}
		if strings.HasPrefix(service.URL, "https://") {
			result = client.PerformHTTPSCheck(service.URL)
		} else if strings.HasPrefix(service.URL, "http://") {
			result = client.PerformHTTPCheck(service.URL)
		} else {
			log.Print("[ERROR][agent] invalid address", service.URL)
		}

		js, _ := json.Marshal(result)
		log.Println(string(js))

		lock.Unlock()

		time.Sleep(5 * time.Second)
	}
}
