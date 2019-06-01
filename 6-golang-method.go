package main

import "fmt"

// GolangMethod : print GolangMethod function in main function
func GolangMethod() {
	car := car{22314, 0, 12562, 225.0}

	fmt.Println("gas_pedal:", car.gasPedal, ",brake_pedal:", car.brakePedal, ",steering_wheel:", car.steeringWheel)
	fmt.Println("speed:", car.mph())
}

const (
	usixteenbitmax float64 = 65535
	kmhMultiple    float64 = 1.60934
)

func (c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / kmhMultiple / usixteenbitmax)
}
