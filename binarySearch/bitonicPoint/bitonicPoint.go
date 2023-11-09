package main

import "fmt"

// Bitonic point is point where the Array start decresing after initial increase
// arr [] = {1,15,25,42,21,17,12,11}
// Bitonic point is 42
// IF the arr is increasing, then print last element of arr
func main() {
	input := []int{1, 45, 47, 50}

	res := bitonicPoint(input)
	fmt.Printf("Bitonic point is %d ", res)
}
func bitonicPoint(input []int) int {
	n := len(input)
	l := 0
	h := n - 1
	for l <= h {
		mid := l + (h-l)/2
		if (mid == 0 || input[mid-1] < input[mid]) && (mid == n-1 || input[mid] > input[mid+1]) {
			return input[mid]
		}
		if mid == 0 || input[mid-1] < input[mid] {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}
	return -1
}
