package main

func main() {

}

// O(1) + O(1) = O(1)
func stupid_function(arr []int) int {
	total := 0   // O(1)
	return total // O(1)
}

func find_sum(arr []int) int {
	total := 0              // O(1)
	for _, v := range arr { // O(1)
		total += v // O(1)
	}

	return total // O(1)
}
