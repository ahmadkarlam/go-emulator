package main

import (
	"log"
)

func main() {
	bus := NewBus()
	cpu := NewCPU(bus)
}

func showAllDataOnBus(bus Bus) {
	for address, value := range bus.ram {
		log.Printf("address: %v, value: %v", address, value)
	}
}
