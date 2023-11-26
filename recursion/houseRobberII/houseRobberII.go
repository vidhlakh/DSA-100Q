package main

import "fmt"

// https://leetcode.com/problems/house-robber-ii

func main() {
	nums := []int{2, 7, 9, 3, 1}

	fmt.Println("House Robber II", rob(nums))

}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	//return max(helperM1(nums, 1, n-1), helperM1(nums, 0, n-2))
	return max(helperM2(nums, n-1, 1), helperM2(nums, n-2, 0))
}

// n-1 to 0
func helperM1(nums []int, i, ind int) int {

	if ind < i {
		return 0
	}
	// rob
	rob := nums[ind] + helperM1(nums, i, ind-2)

	// dont rob
	dontrob := helperM1(nums, i, ind-1)
	return max(rob, dontrob)
}

// 0 to n-1
func helperM2(nums []int, i, ind int) int {

	if ind > i {
		return 0
	}
	// rob
	rob := nums[ind] + helperM2(nums, i, ind+2)

	// dont rob
	dontrob := helperM2(nums, i, ind+1)
	return max(rob, dontrob)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
