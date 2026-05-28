package main

import "fmt"

// A closure is a function that remembers the variables from its surrounding scope even after that scope has finished executing.

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Discount Calculator

func DiscountApplier(discount float64) func(float64) float64 {
	return func(price float64) float64 {
		return price - (price * discount / 100)
	}
}

// Rate Limiter:- API Calls

func rateLimiteer(maxCalls int) func() bool {
	calls := 0
	return func() bool {
		if calls >= maxCalls {
			fmt.Println("Rate limit exceeded")
			return false
		}
		calls++
		fmt.Println("Request check and working")
		return true
	}
}

func main() {
	c := counter()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

	// Off for 10%
	tenPercent := DiscountApplier(10)
	fmt.Println(tenPercent(1000))
	// 900

	// Rate limiter Calls
	apiCalls := rateLimiteer(3)
	apiCalls()

}
