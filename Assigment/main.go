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

	// 1. Create map with 5 contacts
	phoneBook := map[string]string{
		"Alice": "123-456-7890",
		"Bob":   "987-654-3210",
		"Carol": "555-123-4567",
		"David": "444-987-6543",
		"Eve":   "333-111-2222",
	}

	// 2. Print one contact
	fmt.Println("Alice:", phoneBook["Alice"])

	// 3. Safe lookup for missing contact
	val, ok := phoneBook["Zara"]
	if ok {
		fmt.Println("Zara:", val)
	} else {
		fmt.Println("Zara not found")
	}

	// 4. Delete one contact + print count
	delete(phoneBook, "Bob")
	fmt.Println("Total contacts:", len(phoneBook))

	// 5. Loop through all contacts
	for name, number := range phoneBook {
		fmt.Println(name, "→", number)
	}

	// Bonus — delete only if exists
	if _, ok := phoneBook["David"]; ok {
		delete(phoneBook, "David")
		fmt.Println("David deleted")
	}
}
