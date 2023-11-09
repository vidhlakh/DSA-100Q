package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{12, 34, 67, 90}
	N := 4 // num of books
	M := 2 // num of students

	res := bookAllocation(input, N, M)
	fmt.Printf("Minimum value is %d ", res)
}

func bookAllocation(input []int, N, M int) int {
	// search space (maximum num of pages)
	n := len(input)
	l := input[n-1] // as the array is sorted max value is the last value
	var sum int
	res := math.MaxInt
	for _, in := range input {
		sum += in
	}
	h := sum

	for l <= h {
		mid := l + (h-l)/2
		if isAllocationPossible(M, mid, input) {
			res = mid
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}

// Helper function
func isAllocationPossible(M, maxPages int, input []int) bool {
	var currentPages int
	stu := 1

	for _, in := range input {

		if currentPages+in > maxPages {
			stu++
			if stu > M {
				return false
			}
			currentPages = in

		} else {
			currentPages += in
		}
	}
	return true
}
