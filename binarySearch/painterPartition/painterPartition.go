package main

import (
	"fmt"
	"math"
)

// https://practice.geeksforgeeks.org/problems/the-painters-partition-problem1535/1

func main() {
	input := []int{5, 10, 30, 20, 15}
	N := 5 // num of paint boards
	M := 3 // num of painters

	res := painterPartition(input, N, M)
	fmt.Printf("Minimum value is %d ", res)
}

func painterPartition(input []int, NumBoards, NumPainters int) int {
	// search space (maximum num of pages)
	n := len(input)
	l := input[0]
	for i := 1; i < n; i++ {
		l = max(l, input[i])
	}
	var sum int
	res := math.MaxInt
	for _, in := range input {
		sum += in
	}
	h := sum

	for l <= h {
		mid := l + (h-l)/2
		if isAllocationPossible(NumPainters, mid, input) {
			res = mid
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res

}

// Helper function
func isAllocationPossible(NumPainters, maxPages int, input []int) bool {
	var currentPages int
	stu := 1

	for _, in := range input {

		if currentPages+in > maxPages {
			stu++
			if stu > NumPainters {
				return false
			}
			currentPages = in

		} else {
			currentPages += in
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
