package main

import "fmt"

// https://leetcode.com/problems/partition-equal-subset-sum/description/
func main() {
	nums := []int{1, 2, 3, 5}
	fmt.Printf("Recursion:: subset of %v is %t\n", nums, subsetRecursion(nums))

	// Memoization
	fmt.Printf("Memoization::subset of %v is %t\n", nums, subsetMemoization(nums))
	// // Tabulation
	// fmt.Printf("Tabulation::subset of %s,%s is %d\n", nums, subsetTabulation(nums))

	// // Space Optimization
	// fmt.Printf("Space Optimization::subset of %s,%s is %d\n", nums, subsetSpaceOptimization(nums))

}

func subsetRecursion(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2

	return recurse(nums, target, 0)
}

func recurse(nums []int, target, ind int) bool {

	// base condn
	// if target < 0 {
	// 	return false
	// }
	// if target == 0 {
	// 	return true
	// }
	if ind >= len(nums) {
		return target == 0
	}
	// pick
	if recurse(nums, target-nums[ind], ind+1) {
		return true
	}
	// dont pick
	if recurse(nums, target, ind+1) {
		return true
	}

	return false
}

// Memoization

func subsetMemoization(nums []int) bool {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
		for j := 0; j <= target; j++ {
			dp[i][j] = -1
		}
	}
	return memoize(nums, target, 0, dp) == 1
}
func memoize(nums []int, target, i int, dp [][]int) int {
	if target < 0 {
		return 0
	}
	if target == 0 {
		return 1
	}
	if i >= len(nums) {
		return 0
	}
	if dp[i][target] != -1 {
		return dp[i][target]
	}

	// pick
	if memoize(nums, target-nums[i], i+1, dp) == 1 {
		dp[i][target] = 1
		return dp[i][target]
	}
	// dont pick
	if memoize(nums, target, i+1, dp) == 1 {
		dp[i][target] = 1
		return dp[i][target]
	}

	return dp[i][target]
}
