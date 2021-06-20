package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/galenguyer/retina/agent"
)

func main() {
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
