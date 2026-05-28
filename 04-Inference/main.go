package main

import "fmt"

// Inference is the process of automatically determining the type of a variable based on the value assigned to it. In Go, you can use the := operator to declare and initialize a variable without explicitly specifying its type. The Go compiler will infer the type based on the value you provide.
func main() {
	var name = "suhag"
	var age = 24
	var isStudent = true
	fmt.Println("Hello my name is", name, " my age is ", age, "and I am a student?(true/false)", isStudent)

	// Even shorter way to infer
	firstName := "Suhag"
	lastName := "Shetty"
	fmt.Println("First name is ", firstName, " last name is ", lastName)
}
