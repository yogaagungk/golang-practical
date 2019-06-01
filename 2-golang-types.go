package main

import (
	"fmt"
)

func add(x, y float32) float32 {
	return x + y
}

// GolangType : print result of module number 2. Golang Types
func GolangType() {
	// var numberA, numberB float64 = 5.6, 9.5 // with explicit data type, just define
	numberA, numberB := float32(5.6), float32(9.5) // := is operator for assigment

	fmt.Println(add(numberA, numberB))
}
