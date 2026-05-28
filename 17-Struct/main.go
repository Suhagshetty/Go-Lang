package main

import "fmt"

// Structs:- holds multiple data types.

type Student struct {
	Name  string
	Age   int
	Marks float64
}

// REAL WORLD EXAMPLE

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func total(p Product) float64 {
	return p.Price * float64(p.Quantity)
}

func main() {
	student1 := Student{
		Name:  "Suhag",
		Age:   22,
		Marks: 8.23,
	}
	student2 := Student{
		Name:  "Vinay",
		Age:   22,
		Marks: 7.81,
	}
	fmt.Println(student1.Name) // Suhag
	fmt.Println(student2.Age)  // 22

	phone := Product{
		Name:     "Iphone13",
		Price:    69.999,
		Quantity: 1,
	}
	fmt.Println(phone.Name)
	fmt.Println(phone.Price)
}
