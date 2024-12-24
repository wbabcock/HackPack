
package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/irremote"
	"tinygo.org/x/drivers/servo"
)

// ****************************************************************************
// *	Initialization
// ****************************************************************************
func init() {
	// Wait for Serial to be ready (optional for some setups)
	machine.Serial.Configure(machine.UARTConfig{})
	machine.Serial.Write([]byte("Connected" + "\r\n"))

	// Set the pin as OUTPUT to show board is powered ON
	ledPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledPin.High()

	pwm := machine.Timer1
	pwm.Configure(machine.PWMConfig{
		Period: uint64(20 * time.Millisecond), // 20ms period for 50Hz PWM (common for servos)
	})

	yawServo, _ = servo.New(pwm, yawPin)
	pitchServo, _ = servo.New(pwm, pitchPin)
	rollServo, _ = servo.New(pwm, rollPin)

	pinIR.Configure(machine.PinConfig{Mode: machine.PinInput})
	irReciever = irremote.NewReceiver(pinIR)
	irReciever.Configure()

	machine.Serial.Write([]byte("Ready to receive signals" + "\r\n"))
}

// ****************************************************************************
// *	Confirmations
// ****************************************************************************
func nodYes(count int) {
	if count <= 0 {
		count = 3
	}

	startAngle := pitchPOS // Current position of the pitch servo
	//lastAngle := pitchServo
	nodAngle := startAngle + 20 // Angle for nodding motion

	for i := 0; i < count; i++ { // Repeat nodding motion three times
		// Nod up
		for angle := startAngle; angle <= nodAngle; angle++ {
			pitchServo.SetAngle(angle)
			time.Sleep(time.Millisecond * 7) // Adjust delay for smoother motion
		}
		time.Sleep(time.Millisecond * 50) // Pause at nodding position
		// Nod down
		for angle := nodAngle; angle >= startAngle; angle-- {
			pitchServo.SetAngle(angle)
			time.Sleep(time.Millisecond * 7) // Adjust delay for smoother motion
		}
		time.Sleep(time.Millisecond * 50) // Pause at starting position
	}
}

func shakeNo(count int) {
	if count <= 0 {
		count = 3
	}

	for i := 0; i < count; i++ { // Repeat nodding motion three times
		yawServo.SetAngle(140)
		time.Sleep(time.Millisecond * 190)
		yawServo.SetAngle(yawSpeed)
		time.Sleep(time.Millisecond * 50)
		yawServo.SetAngle(40)
		time.Sleep(time.Millisecond * 190)
		yawServo.SetAngle(yawSpeed)
		time.Sleep(time.Millisecond * 50)
	}
}

// ****************************************************************************
// *	Movements
// ****************************************************************************
func moveLeft() {
	// adding the servo speed = 180 (full counterclockwise rotation speed)
	yawServo.SetAngle(yawStop + yawSpeed)
	// stay rotating for a certain number of milliseconds
	time.Sleep(time.Millisecond * yawTime)
	// stop rotating
	yawServo.SetAngle(yawStop)
	// delay for smoothness
	time.Sleep(time.Millisecond * 5)
	// debug note
	machine.Serial.Write([]byte("Move Left" + "\r\n"))
}

func moveRight() {
	// subtracting the servo speed = 0 (full clockwise rotation speed)
	yawServo.SetAngle(yawStop - yawSpeed)
	time.Sleep(time.Millisecond * yawTime)
	yawServo.SetAngle(yawStop)
	time.Sleep(time.Millisecond * 5)
	machine.Serial.Write([]byte("Move Right" + "\r\n"))
}

func lookUp() {
	if pitchPOS > pitchMin { //make sure the servo is within rotation limits (greater than 10 degrees by default)
		pitchPOS = pitchPOS - pitchSpeed //decrement the current angle and update
		pitchServo.SetAngle(pitchPOS)
		time.Sleep(time.Millisecond * 0)
		machine.Serial.Write([]byte("Look Up" + "\r\n"))
	}
}

func lookDown() {
	if pitchPOS < pitchMax { //make sure the servo is within rotation limits (greater than 10 degrees by default)
		pitchPOS = pitchPOS + pitchSpeed //decrement the current angle and update
		pitchServo.SetAngle(pitchPOS)
		time.Sleep(time.Millisecond * 50)
		machine.Serial.Write([]byte("Look Down" + "\r\n"))
	}
}

// ****************************************************************************
// *	Launch
// ****************************************************************************
func fire() {
	rollServo.SetAngle(rollStop + rollSpeed)
	time.Sleep(time.Millisecond * rollTime)
	rollServo.SetAngle(rollStop)
	time.Sleep(time.Millisecond * 5)
	machine.Serial.Write([]byte("Fire Single" + "\r\n"))
}

func fireAll() {
	rollServo.SetAngle(rollStop + rollSpeed)
	time.Sleep(time.Millisecond * rollTime * 6)
	rollServo.SetAngle(rollStop)
	time.Sleep(time.Millisecond * 5)
	machine.Serial.Write([]byte("Fire All" + "\r\n"))
}

// ****************************************************************************
// *	IR Sensor Handler
// ****************************************************************************
func irHandler(data irremote.Data) {
	// Main Loop
	switch irCmdButtons[data.Command] {
	case "left":
		moveLeft()
	case "right":
		moveRight()
	case "up":
		lookUp()
	case "down":
		lookDown()
	case "ok":
		fire()
	case "star":
		fireAll()
	}
}

// ****************************************************************************
// *	Entry point
// ****************************************************************************
func main() {
	irReciever.SetCommandHandler(irHandler)
	for {
		time.Sleep(time.Millisecond * 10)
	}
}
