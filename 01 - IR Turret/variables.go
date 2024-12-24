package main

import (
	"machine"

	"tinygo.org/x/drivers/irremote"
	"tinygo.org/x/drivers/servo"
)

var irCmdButtons = map[uint16]string{
	0x08: "left",
	0x5A: "right",
	0x52: "up",
	0x18: "down",
	0x1C: "ok",
	0x45: "1",
	0x46: "2",
	0x47: "3",
	0x44: "4",
	0x40: "5",
	0x43: "6",
	0x07: "7",
	0x15: "8",
	0x09: "9",
	0x19: "0",
	0x16: "star",
	0x0D: "hash",
}

var (
	irReciever irremote.ReceiverDevice

	// Servos
	pwm        machine.PWM
	yawServo   servo.Servo
	pitchServo servo.Servo
	rollServo  servo.Servo

	// Current values for the three servos
	yawPOS   int
	pitchPOS int = 100
	rollPOS  int
)
