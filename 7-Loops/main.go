package main

import "fmt"

func main() {

	// For Loop:-init → runs once before loop starts
	// condition ; post → checked before each iteration
	//  post    → runs after each iteration

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// To check for even and odd numbers

	nums := []int{2, 3, 4, 5, 6}
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
