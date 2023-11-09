package main

import "fmt"

func main() {
	input := []int{3, 4, 5, 6, 7, 2}

	res := minimumNumberOnSortedRotatedArray(input)
	fmt.Printf("Minimum value is %d ", res)
}

func minimumNumberOnSortedRotatedArray(input []int) int {
	n := len(input)
	l := 0
	h := n - 1
	var mid int
	for l <= h {
		mid = l + (h-l)/2
		if mid == 0 || input[mid] < input[mid-1] {
			return input[mid]
		}
		if input[l] > input[mid] { // unsorted left
			h = mid - 1
		} else if input[mid] > input[h] { // unsorted right
			l = mid + 1
		} else {
			return input[l]
		}
	}
	return -1
}
