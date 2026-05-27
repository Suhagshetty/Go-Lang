package main

import "fmt"

// Constants in Go are values that cannot change after they're set — the compiler will error if you try to reassign them.

// constants declared outside the main function:- they can be accessed anywhere.

const Pi = 3.14159
const myName = "Suhag"
const maxScore = 100

func main() {
	// constant declared inside the main function.
	const gravity = 9.81
	fmt.Println("Name: ", myName)
	fmt.Println("Pi: ", Pi)
	fmt.Println("Max score: ", maxScore)
	// myName = "Kumar" this would cause compile error

	// grouped constants:- are group of constants
	const (
		StatusOk       = 200
		StatusNotFound = 404
		StatusError    = 500
	)
	fmt.Println("Status Ok: ", StatusOk)
	fmt.Println("Status Not Found: ", StatusNotFound)
	fmt.Println("Status Error: ", StatusError)
}
