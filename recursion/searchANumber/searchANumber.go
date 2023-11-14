package main

import "fmt"

func main() {
	fmt.Println("Search a number using recursion")
	input := []int{1, 4, 7, 12, 45, 67, 78, 98, 123, 450, 700}
	k := 45 // Number to find
	N := len(input)

	res := searchANumberUsingRecursion(input, k, 0, N-1)
	if res == 1 {
		fmt.Printf("Number %d is present\n", k)
	} else {
		fmt.Printf("Number %d is not present\n", k)
	}
}

func searchANumberUsingRecursion(input []int, k, l, h int) int {
	if l > h {
		return -1
	}
	mid := l + (h-l)/2
	if input[mid] > k {
		searchANumberUsingRecursion(input, k, l, mid-1)
	} else if input[mid] < k {
		searchANumberUsingRecursion(input, k, mid+1, h)
	}
	return 1

}
