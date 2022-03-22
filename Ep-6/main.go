package main

import "fmt"

func main() {
	fmt.Println("Fact of 4:", fact(4))
}

func fact(n int) int {
	// assuming that n is a positive integer or 0
	if n >= 1 {
		return n * fact(n-1)
	}
	return 1
}
