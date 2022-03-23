package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3}
	fmt.Println(a1)
	quicksort(a1)
	fmt.Println(a1)
}

func quicksort(arr []int) {
	qs(arr, 0, len(arr)-1)
}

func qs(arr []int, l int, r int) {
	if l >= r {
		return
	}
	p := partition(arr, l, r)

	qs(arr, l, p-1)
	qs(arr, p+1, r)
}

func partition(arr []int, l int, r int) int {
	pivot := arr[r]
	i := l - 1
	for _, j := range arr[0 : r-1] {
		if arr[j] < pivot {
			i += 1
			arr[i], arr[j] = arr[j], arr[i]
		}
		arr[i+1], arr[r] = arr[r], arr[i+1]
	}
	return i + 1
}
