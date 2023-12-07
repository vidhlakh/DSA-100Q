package main

import "fmt"

// https://leetcode.com/problems/palindromic-substrings/
func main() {
	s := "abc"
	fmt.Printf("Memoization for palindrome substring %s is %d", s, countSubstrings(s))
}
func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	return memoize(s, n, dp)
}

// memoization
func memoize(s string, n int, dp [][]int) int {
	var res int
	for l := 1; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			if i == j {
				dp[i][j] = 1

			} else if j == i+1 {
				if s[i] == s[j] {
					dp[i][j] = 1 //true

				} else {
					dp[i][j] = 0 // false
				}
			} else if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]

			}

			if dp[i][j] == 1 {
				res++
			}
		}
	}
	return res
}
