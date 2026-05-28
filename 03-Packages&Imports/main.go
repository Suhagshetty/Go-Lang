package main

import (
	"fmt"
	"math"
)

// Round:- rounds to nearest integer.
// Floor:- rounds down to the nearest integer.
// Ceil:- rounds up to the nearest highest integer.
// Trunc:- rounds towards zero to the nearest integer.

func main() {
	fmt.Println("These are some of the math functions")
	printRound(8.7)
	printFloor(7.2)
	printCeil(4.1)
	printTrunc(4.6)
}

func printRound(x float64) {
	fmt.Println("round(", x, "):", math.Round(x))
}
func printFloor(x float64) {
	fmt.Println("Floor(", x, "):", math.Floor(x))
}
func printCeil(x float64) {
	fmt.Println("Ceil(", x, "):", math.Ceil(x))
}
func printTrunc(x float64) {
	fmt.Println("Truncate(", x, "):", math.Trunc(x))
}
