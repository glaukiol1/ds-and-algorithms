package main

import (
	"fmt"
	"time"
)

func main() {
	var array []int
	for s := 0; s < 100*100; s++ {
		array = append(array, s)
	}
	time, total := get_sum(array)
	fmt.Println("get_sum() " + fmt.Sprint(len(array)) + " elements: " + fmt.Sprint(time) + " ms, sum: " + fmt.Sprint(total))
	for s := 0; s < 100*100*100; s++ {
		array = append(array, s)
	}
	time, total = get_sum(array)
	fmt.Println("get_sum() " + fmt.Sprint(len(array)) + " elements: " + fmt.Sprint(time) + " ms, sum: " + fmt.Sprint(total))
	for s := 0; s < 100*100*100*100; s++ {
		array = append(array, s)
	}
	time, total = get_sum(array)
	fmt.Println("get_sum() " + fmt.Sprint(len(array)) + " elements: " + fmt.Sprint(time) + " ms, sum: " + fmt.Sprint(total))
}

func get_sum(n []int) (int64, int) {
	start := time.Now()
	total := 0
	for _, v := range n {
		total += v
	}

	return time.Since(start).Milliseconds(), total
}
