// main.go
package main

import (
	"log"
	"time"

	"github.com/songgao/water"
)

const (
	// Name will be the name of the tunnel
	Name = "tun0"
)

func main() {
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

	time.Sleep(10 * time.Minute)
}
