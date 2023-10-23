package main

import "fmt"

// Binary search works only for sorted elements
func main() {
	fmt.Println("Search a number in log n complexity is done by Binary search")
	input := []int{1, 4, 7, 12, 45, 67, 78, 98, 123, 450, 700}
	k := 45 // Number to find

	res := searchANumber(input, k)
	if res == 1 {
		fmt.Printf("Number %d is present\n", k)
	} else {
		fmt.Printf("Number %d is not present\n", k)
	}
}

// Time complexity: O(log n) Space complexity: O(1)
func searchANumber(input []int, k int) int {

	low := 0
	high := len(input) - 1

	for low <= high {
		mid := low + (high-low)/2 // to avoid overflow, say if low,high are intmax, low + high is twice the intmax which overflows

		if input[mid] > k {
			fmt.Println("camhe", input[mid], mid)
			high = mid - 1
		} else if input[mid] < k {
			fmt.Println("camlow", input[mid], mid)
			low = mid + 1
		} else {
			return 1
		}

	}

	return -1
}
