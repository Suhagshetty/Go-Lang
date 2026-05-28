package main

import "fmt"

func main() {
	// range is used to iterate over arrays, slices, maps, strings, and channels. It's Go's clean alternative to a traditional for loop with an index counter.

	cart := map[string]float64{
		"Nike Shoes": 4999.99,
		"T-shirt":    999.98,
		"Jeans":      4888.99,
		"Watch":      1799.88,
		"Sunglass":   2222.88,
	}
	// Calculate total
	total := 0.0
	for _, price := range cart {
		total += price
	}
	fmt.Println("The total price is ", total)

	// To find the most expensive item in the cart

	expensive := ""
	highestPrice := 0.0
	for item, price := range cart {
		if price > highestPrice {
			highestPrice = price
			expensive = item
		}
	}
	fmt.Println("The most expensive item is ", expensive, " with a price of ", highestPrice)
}
