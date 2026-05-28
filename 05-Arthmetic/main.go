package main

import "fmt"

func main() {
	a := 10
	b := 10

	// Basic Arithmetic Operations

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// Float Division

	c := 10.0
	d := 3.0
	fmt.Println("Floating division:- ", c/d)

	// Short hand operators

	x := 10
	x += 10
	fmt.Println("x after += 10: ", x)
}
