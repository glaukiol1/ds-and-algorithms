package main

import (
	"fmt"
)

func main() {
	arr := []float64{}
	for s := 0; s < 100e16; s = s + 1 {
		arr = append(arr, float64(s))
	}
	search := 100e16
	fmt.Println("(Linear) Indx of 11: " + fmt.Sprint(
		linear_search(arr, search)))

	fmt.Println("(Binary) Indx of 11: " + fmt.Sprint(
		binary_search(arr, search)))
}

func linear_search(arr []float64, target float64) int {
	for i, v := range arr {
		if target == v {
			return i
		}
	}
	return -1
}

func binary_search(arr []float64, target float64) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if target < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
