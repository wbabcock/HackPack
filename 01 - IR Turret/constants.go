package main

import (
	"machine"
)

const (
	// On-Board LED to show powered on
	ledPin = machine.LED

	// IR Pin
	pinIR = machine.D9

	// Servo Pins
	yawPin   = machine.D10
	pitchPin = machine.D11
	rollPin  = machine.D12

	// Servo Settings
	pitchSpeed = 8
	pitchMax   = 175 // max is 180...keep below
	pitchMin   = 10  // min is 0... keep above

	yawSpeed = 90
	yawStop  = 90  // keep 90 for smooth and accurate movement
	yawTime  = 150 // in milliseconds

	rollSpeed = 90
	rollStop  = 90  // keep 90 for smooth and accurate movement
	rollTime  = 150 // in milliseconds
)
