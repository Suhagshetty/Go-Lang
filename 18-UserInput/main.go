package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Marks float64
}

func main() {
	var s Student
	fmt.Println("Enter the name:- ")
	fmt.Scan(&s.Name)
	fmt.Println("Enter the age:- ")
	fmt.Scan(&s.Age)
	fmt.Println("Enter the marks:- ")
	fmt.Scan(&s.Marks)

	fmt.Println("\n----Student Details----")
	fmt.Println("Name: ", s.Name)
	fmt.Println("Age: ", s.Age)
	fmt.Println("Marks: ", s.Marks)
}
