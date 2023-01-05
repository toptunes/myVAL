package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/songgao/water"
)

const (
	// Name will be the name of the tunnel
	Name = "tun0"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		done <- true
	}()

	config := water.Config{
		DeviceType: water.TUN,
	}
	config.Name = Name

	// Create a tunnel
	inf, err := water.New(config)
	if err != nil {
		log.Fatalf("error while creating a tun interface: %v", err)
	}
	defer inf.Close()

	log.Printf("tunnel created with name: %s", inf.Name())
	<-done
	log.Println("Exiting....")
}
