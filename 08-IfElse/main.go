package main

import (
	"fmt"
)

// If Else statement in Go

func main() {
	age := 17
	if age >= 18 {
		fmt.Println("you are above 18 so u can drink")
	} else {
		fmt.Println("U are under age MF")
	}

	const AgeForDriverLicense = 18
	if AgeForDriverLicense >= 18 {
		fmt.Println("You are eligible for DL")
	} else {
		fmt.Println("Not eligible for DL")
	}

	i := 0
	for i = 0; i <= 20; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}

}
