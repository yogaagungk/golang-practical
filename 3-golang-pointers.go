package main

import (
	"fmt"
)

// GolangPointers : print result of module number 3. Golang Pointers
func GolangPointers() {
	x := 15
	a := &x // a hold memory address of x

	fmt.Println(a)  // print a memory address
	fmt.Println(*a) // print a value through the pointer

	*a = 5 // re assign value of pointer, also change value of x

	fmt.Println(x)

	*a = *a * *a // re assign value of pointer

	fmt.Println(x)  // print new value
	fmt.Println(*a) // print new value

	fmt.Println(&x) // print memory address
	fmt.Println(a)  // print memory address

}
