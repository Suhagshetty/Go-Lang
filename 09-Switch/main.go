package main

import "fmt"

func main() {
	// The basic switch statement

	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week lets go!")
	case "Tuesday":
		fmt.Println("Second day of the weekend plan stuff!")
	case "Saturday":
		fmt.Println("Its the weekend")
	default:
		fmt.Println("Just a normal day")
	}

	// Switch with no variables

	age := 20
	switch {
	case age < 13:
		fmt.Println("Child")
	case age < 18:
		fmt.Println("Teenager")
	case age >= 18:
		fmt.Println("Adult")
	}

	grade := 85
	switch {
	case grade >= 90:
		fmt.Println("A")
	case grade >= 80:
		fmt.Println("B")
	case grade >= 70:
		fmt.Println("C")
	default:
		fmt.Println("JUST PASS")
	}

	// Multiple values in one case

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("weekend")
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	}

	role := "Admin"
	switch role {
	case "Admin":
		fmt.Println("Full access")
	case "User":
		fmt.Println("Limited access")
	}
}
