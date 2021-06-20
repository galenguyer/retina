package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/galenguyer/retina/agent"
	"github.com/galenguyer/retina/config"
	"gopkg.in/yaml.v2"
)

func main() {
	conf := loadConfig()
	_ = conf

	agent.Start()
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

func loadConfig() *config.Config {
	customConfigFile := os.Getenv("RETINA_CONFIG_FILE")
	if len(customConfigFile) > 0 {
	} else {
		conf, err := config.Load("./config.yaml")
		if err != nil {
			log.Fatal(err)
		}
		y, _ := yaml.Marshal(conf)
		fmt.Println(string(y))
	}
	return nil
}
