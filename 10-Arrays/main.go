package main

import "fmt"

func main() {
	var nums [5]int
	fmt.Println(nums)

	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	fmt.Println(numbers)
	// [10 20 0 0 0]

	fmt.Println(len(numbers)) // 5

	// String arrays
	names := [3]string{"Suhag", "Sumana", "Santhosh"}
	fmt.Println(len(names)) // 3
	fmt.Println(names[0])   // Suhag

	// FIND THE LARGEST NUMBER IN ARRAY

	numberss := [5]int{10, 20, 30, 40, 50}
	largest := numberss[0]
	for _, value := range numberss {
		if value > largest {
			largest = value
		}
	}
	fmt.Println(largest)
}
