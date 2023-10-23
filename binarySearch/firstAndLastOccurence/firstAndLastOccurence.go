package main

import (
	"fmt"
	"sort"
)

// Binary search works only for sorted elements
// Input is sorted numbers with  consecutive repeated elements
func main() {
	fmt.Println("Search a numberFind first and last occurence of a number  in log n complexity is done by Binary search")
	input := []int{1, 4, 7, 12, 45, 46, 46, 46, 46, 67, 78, 98, 123, 450, 700}
	k := 46 // Number to find
	var res []int
	res = append(res, Occurence(input, k, true))
	res = append(res, Occurence(input, k, false))
	fmt.Println("First and last Occurence ", res)

	// Using library
	var res2 []int
	res2 = append(res2, binarySearch(input, k, true))
	res2 = append(res2, binarySearch(input, k, false))
	fmt.Println("First and last Occurence ", res2)
}

// Time complexity: O(log n) Space complexity: O(1)
func Occurence(input []int, k int, first bool) int {

	low := 0
	high := len(input) - 1
	occurenceIndex := -1
	for low <= high {
		mid := low + (high-low)/2 // to avoid overflow, say if low,high are intmax, low + high is twice the intmax which overflows

		if input[mid] > k {
			fmt.Println("camhe", input[mid], mid)
			high = mid - 1
		} else if input[mid] < k {
			fmt.Println("camlow", input[mid], mid)
			low = mid + 1
		} else {
			if first {
				high = mid - 1
			} else {
				low = mid + 1
			}
			occurenceIndex = mid
		}
	}

	return occurenceIndex
}

// Using go library >= will return the first occurence of the elemnt
func binarySearch(input []int, k int, first bool) int {
	i := sort.Search(len(input), func(i int) bool {
		if first {
			return input[i] >= k
		}
		return input[i] > k
	})
	if i < len(input) {
		if first {
			return i
		}
		return i - 1
	}
	return -1
}

// // Time complexity: O(log n) Space complexity: O(1)
// func lastOccurence(input []int, k int) int {

// 	low := 0
// 	high := len(input) - 1
// 	last := -1
// 	for low <= high {
// 		mid := low + (high-low)/2 // to avoid overflow, say if low,high are intmax, low + high is twice the intmax which overflows

// 		if input[mid] > k {
// 			fmt.Println("camhe", input[mid], mid)
// 			high = mid - 1
// 		} else if input[mid] < k {
// 			fmt.Println("camlow", input[mid], mid)
// 			low = mid + 1
// 		} else {
// 			last = mid
// 			low = mid + 1
// 		}

// 	}
// 	return last
// }
