package main

import "fmt"

// Pointers:- stores memory address of a variable instead of the value itself. There are 2 operators.

// & -> address of -> gives you the memory address of a Variable.
// * -> gives you the value at that address.

func main() {
	name := "Suhag"
	p := &name
	fmt.Println(p)
	//  0xc000014080  (memory address)
	fmt.Println(*p)
	// Suhag is the value at the address

	*p = "Shetty"
	fmt.Println(name)
	// Suhag Shetty

	age := 21
	j := &age
	fmt.Println(j)
	// 0x140000100d8
	fmt.Println(*j)
	// 21

	*j = 22
	fmt.Println(*j)
	// 2
}
