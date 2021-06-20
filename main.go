package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/galenguyer/retina/agent"
	"github.com/galenguyer/retina/config"
	"github.com/galenguyer/retina/storage"
	"gopkg.in/yaml.v2"
)

func main() {
	conf, _ := loadConfig()

	storage.CreateDatabase()

	go startServer()
	agent.Start(conf)

	signalChannel := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChannel
		log.Println("Received termination signal, attempting to gracefully shut down")
		done <- true
	}()
	<-done
	log.Println("Shutting down")
}

func loadConfig() (conf *config.Config, err error) {
	customConfigFile := os.Getenv("RETINA_CONFIG_FILE")
	if len(customConfigFile) > 0 {
		conf, err = config.Load(customConfigFile)
	} else {
		conf, err = config.Load("./config.yaml")
	}
	if err != nil {
		log.Fatal(err)
	}
	y, _ := yaml.Marshal(conf)
	fmt.Println(string(y))
	return conf, nil
}

func startServer() {
	fileServer := http.FileServer(http.Dir("./web/static/"))
	http.HandleFunc("/api/v1/ok", ok)
	http.Handle("/", http.StripPrefix(strings.TrimRight("/", "/"), fileServer))
	log.Println("starting webserver on port 8000")
	http.ListenAndServe(":8000", nil)
}

type Okt struct {
	Status int
}

func ok(w http.ResponseWriter, r *http.Request) {
	res := &Okt{Status: 200}
	json.NewEncoder(w).Encode(res)
}
