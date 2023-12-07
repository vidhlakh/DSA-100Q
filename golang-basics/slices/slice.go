package main

import "fmt"

func addNum(nums *[][]int, newNum []int) {
	*nums = append(*nums, newNum)
	fmt.Printf("addNum nums %v, %p\n", *nums, nums)
	newNum = []int{8, 9}
	*nums = append(*nums, newNum)
}

func main() {
	var nums [][]int
	newNum := []int{4, 5}
	fmt.Printf("  main nums %v,%p\n", nums, nums)
	addNum(&nums, newNum)
	fmt.Printf("  main nums %v\n", nums)
	addNum(&nums, newNum)
	fmt.Printf("  main nums %v\n", nums)
}
