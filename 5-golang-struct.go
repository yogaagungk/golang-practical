package main

import "fmt"

type car struct {
	gasPedal      uint16
	brakePedal    uint16
	steeringWheel int16
	topSpeedKmh   float64
}

// GolangStruct : print GolangStruct function in main function
func GolangStruct() {
	car := car{22314, 0, 12562, 225.0}

	fmt.Println("gas_pedal:", car.gasPedal, ",brake_pedal:", car.brakePedal, ",steering_wheel:", car.steeringWheel)
}
