package main

import "fmt"

func main() {
	// map [student Name] map [subject] score
	reportCard := map[string]map[string]int{
		"Alice": {
			"Math":    92,
			"Science": 88,
			"English": 76,
		},

		"Bob": {
			"Math":    74,
			"Science": 91,
			"English": 83,
		},
	}
	// Read a specific score
	fmt.Println(reportCard["Alice"])
	// map[English:76 Math:92 Science:88]

	// add a new student
	reportCard["Charlie"] = map[string]int{
		"Math":    66,
		"Science": 54,
		"English": 88,
	}
	fmt.Println(reportCard["Charlie"])
	// map[English:88 Math:66 Science:54]

	// Check if user exists

	val, ok := reportCard["David"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("David is not found")
	}
}
