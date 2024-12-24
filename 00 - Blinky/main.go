package main

import (
	"machine"
	"time"
)

const (
	ledPin = machine.LED // Set the pin to the on-board LED
)

func main() {
	// Set the pin as OUTPUT
	ledPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Loop
	for {
		ledPin.High()           // LED on
		time.Sleep(time.Second) // wait 1 second
		ledPin.Low()            // LED off
		time.Sleep(time.Second) // wait 1 second
	}
}
