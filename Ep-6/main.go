package main

import "fmt"

func main() {
	fmt.Println("Fact of 4:", fact(4))
	fmt.Println("First 8 fibonacci:", fib(8))
}

func fact(n int) int {
	// assuming that n is a positive integer or 0
	if n >= 1 {
		return n * fact(n-1)
	}
	return 1
}

func fib(n int) int {
	if n >= 3 {
		return (fib(n-1) + fib(n-2))
	}
	return 1
}
