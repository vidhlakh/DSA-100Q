package main

import "fmt"

// https://leetcode.com/problems/house-robber/
func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println("House Robber problem:", helper(len(nums)-1, nums))
}

func helper(ind int, nums []int) int {

	if ind < 0 {
		return 0
	}
	// rob
	rob := nums[ind] + helper(ind-2, nums)
	//dontrob
	dontrob := helper(ind-1, nums)
	return max(rob, dontrob)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
