package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/galenguyer/retina/agent"
	"github.com/galenguyer/retina/config"
	"github.com/galenguyer/retina/controller"
	"github.com/galenguyer/retina/storage"
	"gopkg.in/yaml.v2"
)

func main() {
	conf, _ := loadConfig()

	storage.CreateDatabase()

	go controller.StartServer()
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
