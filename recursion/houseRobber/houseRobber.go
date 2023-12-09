package main

import "fmt"

// https://leetcode.com/problems/house-robber/
func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println("House Robber problem:", helper(len(nums)-1, nums))
	// memoize
	fmt.Println("Memoization,House Robber problem:", rob(nums))

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

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}
	return memoize(n, 0, nums, dp)
}

func memoize(n, ind int, nums, dp []int) int {

	if ind >= n {
		return 0
	}
	if dp[ind] != -1 {
		return dp[ind]
	}

	// rob
	rob := nums[ind] + memoize(n, ind+2, nums, dp)
	//dontrob
	dontrob := memoize(n, ind+1, nums, dp)
	dp[ind] = max(rob, dontrob)
	return dp[ind]
}
