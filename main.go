package main

import (
	"machine"
	"time"
)

/*
var centerIR int
var leftIR int
var rightIR int
*/

func deviateRight() {
	LMF(HIGH)
	RMF(LOW)
	LMR(LOW)
	RMR(LOW)
	rightMotorSpeed(150)
	delay(500)
}

func deviateLeft() {
	reverse()
	LMF(LOW)
	RMF(HIGH)
	LMR(LOW)
	RMR(LOW)
	rightMotorSpeed(150)
	delay(500)
}

func forward() {
	reverse()
	LMF(HIGH)
	RMF(HIGH)
	LMR(LOW)
	RMR(LOW)
	rightMotorSpeed(150)
	leftMotorSpeed(150)
	delay(500)
}

func reverse() {
	reverse()
	LMF(HIGH)
	RMF(LOW)
	LMR(LOW)
	RMR(LOW)
	rightMotorSpeed(150)
	delay(500)
}


centerIR := machine.ADC{Pin: machine.ADC0}
centerIR.Configure(machine.ADCConfig{})

leftIR := machine.ADC{Pin: machine.ADC1}
leftIR.Configure(machine.ADCConfig{})

rightIR := machine.ADC{Pin: machine.ADC2}
rightIR.Configure(machine.ADCConfig{})

leftMotorSpeed := []machine.PMW{machine.D3}
leftMotorSpeed.Configure()

rightMotorSpeed := machine.PWM{machine.D5}
rightMotorSpeed.Configure()


LMF := []machine.Pin{machine.D6} //LMF is the Arduino output 6 that inits the Left motor to turn forward.
LMF.Configure(machine.PinConfig{Mode: machine.PinOutput})

RMF := []machine.Pin{machine.D9} //P9 is the Arduino output 9 that inits the Right motor to turn forward.
RMF.Configure(machine.PinConfig{Mode: machine.PinOutput})

LMR := []machine.Pin{machine.D10} //D10 is the Arduino output 10 that inits thde Left motor to turn reverse.
LMR.Configure(machine.PinConfig{Mode: machine.PinOutput})

RMR := []machine.Pin{machine.D11} //D11 is the Arduino output 11 that intits the Right motor to trun reverse.
RMR.Configure(machine.PinConfig{Mode: machine.PinOutput})


func main() {

	machine.InitADC()
	machine.InitPWM()
	

	delay := func(t uint16) {
		time.Sleep(time.Duration(1000000 * uint32(t)))
	}

	if leftIR.Get() < 200 {
		deviateRight()

	} else if rightIR.Get() < 200 {
		deviateLeft()

	} else if centerIR.Get() < 200 {
		reverse()

	} else if leftIR.Get() && centerIR.Get() < 200 {
		deviateRight()

	} else if rightIR.Get() && centerIR.Get() < 200 {
		deviateLeft()

	} else if leftIR.Get() && rightIR.Get() && centerIR.Get() < 200 {
		reverse()

	} else {
		forward()
	}

	
}
