package main

import "fmt"

func main() {
	input := []int{1, 3, 4, 7, 3}
	res := findPeakElement(input)
	fmt.Printf("Peak element is %d", res)

}

func findPeakElement(input []int) int {

	l := 0
	h := len(input) - 1
	var mid int
	for l <= h {
		mid = l + (h-l)/2
		if (mid == 0 || input[mid-1] <= input[mid]) && (mid == len(input)-1 || input[mid+1] <= input[mid]) {
			return mid
		}
		if mid > 0 && input[mid-1] > input[mid] {
			h = mid - 1

		} else {
			l = mid + 1

		}
	}
	return mid
}
