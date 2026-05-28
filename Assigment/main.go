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

	scores := make([]int, 0, 4)
	scores = append(scores, 85, 92, 78, 95)
	fmt.Println(len(scores))
	fmt.Println(cap(scores))

	scores = append(scores, 88)
	fmt.Println(len(scores))
	fmt.Println(cap(scores))

	slicedScores := scores[1:4]
	fmt.Println(slicedScores)

	slicedScores[0] = 999
	fmt.Println(scores)
	fmt.Println(slicedScores)
}
