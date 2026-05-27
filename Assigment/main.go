package main

import (
	"fmt"
	"math"
)

func main() {
	var name = "suhag"
	const schoolName = "SAIT"
	mark1 := 85
	mark2 := 91
	mark3 := 89
	totalMarks := mark1 + mark2 + mark3
	percentage := math.Round((float64(totalMarks) / 300) * 100)
	fmt.Println("School: ", schoolName)
	fmt.Println("Student: ", name)
	fmt.Println("Marks: ", mark1, mark2, mark3)
	fmt.Println("Total: ", totalMarks, "/300")
	fmt.Println("Percentage: ", percentage)
}
