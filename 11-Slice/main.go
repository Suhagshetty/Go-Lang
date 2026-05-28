package main

import "fmt"

func main() {
	// Slice:-A slice is a dynamically-sized, flexible view into the elements of an array. Unlike arrays (which have fixed size), slices can grow and shrink. They're one of Go's most used data structures.

	// Create a slice using a literal

	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println((cap(fruits)))

	// Append to a slice
	fruits = append(fruits, "mango")
	fmt.Println(fruits)
	// [apple banana cherry mango]
	fmt.Println(len(fruits)) // 4
	fruits = append(fruits, "melons")
	fmt.Println(fruits)

	// Slice a slice (slicing)

	slices := fruits[1:3]
	// from index 1 to index 3 but 1 is not included
	fmt.Println(slices)
	// banana cherry
	slices = fruits[1:4]
	fmt.Println(slices)
	// banana cherry mango

	// MAKE()= creates a slice with a specified length and capacity.

	nums := make([]int, 3, 5) // len = 3 cap = 5 (capacity)
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	fmt.Println(nums)
	fmt.Println(cap(nums))
	nums = append(nums, 40)
	nums = append(nums, 50)
	fmt.Println((nums))

	/* What cap actually means:-

		make([]int, 3, 5)
		Underlying array:  [ 10 | 20 | 30 |  _ |  _ ]
	                     ^              ^         ^
	                   index 0        len=3     cap=5




	*/

}
