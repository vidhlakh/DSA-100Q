package main

import "fmt"

// https://leetcode.com/problems/longest-palindromic-subsequence/description/
func main() {
	s := "bbbab"
	fmt.Printf("Memoization for longest palindrome substring %s is %d", s, longestPalindromeSubseq(s))
}
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 0
		}
	}
	return memoize(s, n, dp)
}

// memoization
func memoize(s string, n int, dp [][]int) int {

	for l := 1; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			if i == j {
				dp[i][j] = 1
			} else if j == i+1 {
				if s[i] == s[j] {
					dp[i][j] = 2
				} else {
					dp[i][j] = 1
				}
			} else if s[i] == s[j] {
				dp[i][j] = 2 + dp[i+1][j-1]

			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
