package main

import "fmt"

func main() {
	input := []int{5, 6, 7, 8, 9, 10, 1, 2, 3}
	k := 1
	res := searchANumberOnSortedRotatedArray(input, k)
	fmt.Printf("Index of %d is %d ", k, res)
}

func searchANumberOnSortedRotatedArray(input []int, k int) int {

	l := 0
	h := len(input) - 1
	for l <= h {
		mid := l + (h-l)/2
		if input[mid] == k {
			return mid
		}
		if input[l] < input[mid] { // left side is sorted
			if input[l] <= k && k < input[mid] {
				h = mid - 1 // reduce to left side

			} else {
				l = mid + 1
			}
		} else { // right side is sorted
			if k > input[mid] && k < input[h] {
				l = mid + 1
			} else {
				h = mid - 1
			}
		}
	}
	return -1
}
