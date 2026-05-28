package main

import "fmt"

// Function in Go:- A function is a reusable block of code that takes inputs, does something, and optionally returns a result.

/* Syntax:-
func functionName(param1 type, param2 type) returnType {
    // body
    return value
}
*/

func greet(name string) string {
	return "Hello, " + name
}

func evenNumber(number int) bool {
	return number%2 == 0
}

// Function checks even numbers in an array.

func isEvenCheck(n int) bool {
	return n%2 == 0
}
func checkAllEven(nums []int) {
	for _, n := range nums {
		if isEvenCheck(n) {
			fmt.Println(n, "is even")
		} else {
			fmt.Println(n, "is odd")
		}
	}
}

func invite(name string, age int) string {
	return "Hello, " + name + " i am " + fmt.Sprint(age) + " years old"
}

func main() {
	fmt.Println(greet("Suhag")) // Hello, Suhag
	fmt.Println(evenNumber(6))  // True
	fmt.Println(evenNumber(7))  // False
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	checkAllEven(nums)
	fmt.Println(invite("Suhag", 22))
}
