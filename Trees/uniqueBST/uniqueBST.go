package main

import "fmt"

// https://leetcode.com/problems/unique-binary-search-trees/
// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	fmt.Println("Number of unique ways of forming trees", numTrees(4))
}

func numTrees(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = 0
	}
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for root := 1; root <= i; root++ {
			dp[i] += dp[root-1] * dp[i-root]
		}
	}

	return dp[n]
}
