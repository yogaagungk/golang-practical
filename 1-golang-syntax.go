package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func foo() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
}

// GolangSyntax : print result of module number 1. Golang Syntax
func GolangSyntax() {
	foo()

	fmt.Println("It is currently:", time.Now())
	fmt.Println("A number between 1-100:", rand.Intn(100))
}
