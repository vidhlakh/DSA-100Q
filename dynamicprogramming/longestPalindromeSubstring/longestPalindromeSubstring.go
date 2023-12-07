package main

import "fmt"

// https://leetcode.com/problems/longest-palindromic-substring/description/
func main() {
	s := "babad"
	fmt.Printf("Memoization for longest palindrome substring %s is %s", s, longestPalindrome(s))
}

func longestPalindrome(s string) string {
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
func memoize(s string, n int, dp [][]int) string {
	var res string
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
				res = string(s[i : j+1])
			}
		}
	}
	return res
}
