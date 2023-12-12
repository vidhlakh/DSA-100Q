package main

import "fmt"

// https://leetcode.com/problems/longest-common-subsequence/
func main() {
	text1 := "sea"
	text2 := ""
	fmt.Printf("Recursion:: LCS of %s,%s is %d\n", text1, text2, lcsRecursion(text1, text2))

	// Memoization
	fmt.Printf("Memoization::LCS of %s,%s is %d\n", text1, text2, lcsMemoization(text1, text2))
	// Tabulation
	fmt.Printf("Tabulation::LCS of %s,%s is %d\n", text1, text2, lcsTabulation(text1, text2))

	// Space Optimization
	fmt.Printf("Space Optimization::LCS of %s,%s is %d\n", text1, text2, lcsSpaceOptimization(text1, text2))

	fmt.Println("Min delelte sum:", minimumDeleteSum(text1, text2))
}

// Recursion Time : O(2^n+m) Space: O(n+m)+ recursive stack
func lcsRecursion(t1, t2 string) int {
	return recurse(t1, t2, 0, 0)
}

func recurse(t1, t2 string, i1, i2 int) int {
	if i1 == len(t1) || i2 == len(t2) {
		return 0
	}
	// chars are same
	if t1[i1] == t2[i2] {
		return 1 + recurse(t1, t2, i1+1, i2+1)
	}

	// chars are diff
	return max(recurse(t1, t2, i1+1, i2), recurse(t1, t2, i1, i2+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Memoization Time : O(n+m) Space: O(n+m)+ recursive stack
func lcsMemoization(t1, t2 string) int {
	n := len(t1)
	m := len(t2)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = -1
		}
	}
	return memoize(t1, t2, 0, 0, dp)
}

func memoize(t1, t2 string, i, j int, dp [][]int) int {
	if i == len(t1) || j == len(t2) {
		return 0
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	// chars are same
	if t1[i] == t2[j] {
		dp[i][j] = 1 + memoize(t1, t2, i+1, j+1, dp)
		return dp[i][j]
	}

	// chars are diff
	dp[i][j] = max(memoize(t1, t2, i+1, j, dp), memoize(t1, t2, i, j+1, dp))
	return dp[i][j]
}

// Tabulation Time : O(n+m) Space: O(n+m) Iterative (so no recursive stack)
func lcsTabulation(t1, t2 string) int {
	n := len(t1)
	m := len(t2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = 0
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = 0
	}
	return tabulize(t1, t2, n, m, dp)
}
func tabulize(t1, t2 string, n, m int, dp [][]int) int {

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {

			// chars are same
			if t1[i-1] == t2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else { // chars are diff
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
			fmt.Println("i,j", i, j, dp[i][j])
		}
	}
	return dp[n][m]
}

// Space Optimization Time : O(nm) Space: O(m)
func lcsSpaceOptimization(text1, text2 string) int {
	n := len(text1)
	m := len(text2)
	curr := make([]int, m+1)
	prev := make([]int, m+1)
	for i := 0; i <= m; i++ {
		curr[i] = 0
		prev[i] = 0
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {

			// chars are same
			if text1[i-1] == text2[j-1] {
				curr[j] = 1 + prev[j-1]
			} else {
				// chars are diff
				curr[j] = max(curr[j-1], prev[j])
			}
			fmt.Println("i,j", i, j, curr[j])
		}
		temp := make([]int, m+1)
		copy(temp, curr)
		prev = temp
	}
	return curr[m]
}

func minimumDeleteSum(s1 string, s2 string) int {

	n := len(s1)
	m := len(s2)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = -1
		}
	}
	var res int
	for i := 0; i < n; i++ {
		res = int(s1[i])
	}
	for j := 0; j < m; j++ {
		res += int(s2[j])
	}
	return res - 2*(memoizeDel(s1, s2, 0, 0, dp))

}

func memoizeDel(t1, t2 string, i, j int, dp [][]int) int {
	if i == len(t1) || j == len(t2) {
		return 0
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	// chars are same
	if t1[i] == t2[j] {
		fmt.Println("tt", int(t1[i]))
		dp[i][j] = int(t1[i]) + int(t2[j]) + memoizeDel(t1, t2, i+1, j+1, dp)
		return dp[i][j]
	}

	// chars are diff
	dp[i][j] = max(memoizeDel(t1, t2, i+1, j, dp), memoizeDel(t1, t2, i, j+1, dp))
	return dp[i][j]
}
