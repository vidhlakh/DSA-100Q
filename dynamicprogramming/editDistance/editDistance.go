package main

import "fmt"

// https://leetcode.com/problems/edit-distance/submissions/
func main() {
	text1 := "horse"
	text2 := "ros"
	fmt.Printf("Recursion:: editDist of %s,%s is %d\n", text1, text2, editDistRecursion(text1, text2))

	// Memoization
	fmt.Printf("Memoization::editDist of %s,%s is %d\n", text1, text2, editDistMemoization(text1, text2))
	// Tabulation
	fmt.Printf("Tabulation::editDist of %s,%s is %d\n", text1, text2, editDistTabulation(text1, text2))

	// Space Optimization
	fmt.Printf("Space Optimization::editDist of %s,%s is %d\n", text1, text2, editDistSpaceOptimization(text1, text2))

}

// Recursion Time : O(3^n+m) Space: O(n+m)+ recursive stack
func editDistRecursion(t1, t2 string) int {
	return recurse(t1, t2, 0, 0)
}

func recurse(t1, t2 string, i1, i2 int) int {
	if i1 == len(t1) {
		return len(t2) - i2
	}
	if i2 == len(t2) {
		return len(t1) - i1
	}
	// chars are same
	if t1[i1] == t2[i2] {
		return recurse(t1, t2, i1+1, i2+1)
	}

	// chars are diff
	// Replace i1+1,i2+1
	// Insert i1+1,i2
	// Delete i1,i2+1

	return 1 + min(recurse(t1, t2, i1+1, i2), recurse(t1, t2, i1, i2+1), recurse(t1, t2, i1+1, i2+1))
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < a && b < c {
		return b
	}
	return c
}

// Memoization Time : O(n+m) Space: O(n+m)+ recursive stack
func editDistMemoization(t1, t2 string) int {
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
	if i >= len(t1) {
		return len(t2) - j
	}
	if j >= len(t2) {
		return len(t1) - i
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	// chars are same
	if t1[i] == t2[j] {
		dp[i][j] = memoize(t1, t2, i+1, j+1, dp)
		return dp[i][j]
	}

	// chars are diff
	dp[i][j] = 1 + min(memoize(t1, t2, i+1, j, dp), memoize(t1, t2, i, j+1, dp), memoize(t1, t2, i+1, j+1, dp))

	return dp[i][j]
}

// Tabulation Time : O(nm) Space: O(nm) Iterative (so no recursive stack)
func editDistTabulation(t1, t2 string) int {
	n := len(t1)
	m := len(t2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}
	return tabulize(t1, t2, n, m, dp)
}
func tabulize(t1, t2 string, n, m int, dp [][]int) int {

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {

			// chars are same
			if t1[i-1] == t2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else { // chars are diff
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
			fmt.Println("i,j", i, j, dp[i][j])
		}
	}
	return dp[n][m]
}

// Space Optimization Time : O(nm) Space: O(m)
func editDistSpaceOptimization(t1, t2 string) int {
	n := len(t1)
	m := len(t2)
	if n == 0 {
		return m
	}
	curr := make([]int, m+1)
	prev := make([]int, m+1)

	for j := 0; j <= m; j++ {
		prev[j] = j
	}
	return optimize(t1, t2, n, m, curr, prev)
}

func optimize(t1, t2 string, n, m int, curr, prev []int) int {

	for i := 1; i <= n; i++ {
		curr[0] = i
		for j := 1; j <= m; j++ {

			// chars are same
			if t1[i-1] == t2[j-1] {
				curr[j] = prev[j-1]
			} else {
				// chars are diff
				curr[j] = 1 + min(curr[j-1], prev[j], prev[j-1])
			}
			//fmt.Println("i,j", i, j, curr[j])
		}
		temp := make([]int, m+1)
		copy(temp, curr)
		prev = temp
	}
	return curr[m]
}
