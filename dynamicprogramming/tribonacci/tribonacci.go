package main

import "fmt"

func main() {
	n := 5
	fmt.Printf("Recursion tribonacci for  %d is %d\n", n, recurseTribonacci(n))
	fmt.Printf("Memoization tribonacci for  %d is %d\n", n, memoizeTribonacci(n))
	fmt.Printf("Bottom up tribonacci for  %d is %d\n", n, tabulizeTribonacci(n))
	fmt.Printf("Space optimize tribonacci for  %d is %d\n", n, optimizeTribonacci(n))

}

func recurseTribonacci(n int) int {
	if n <= 1 {
		return n
	}
	if n == 2 {
		return 1
	}
	return recurseTribonacci(n-1) + recurseTribonacci(n-2) + recurseTribonacci(n-3)
}

func memoizeTribonacci(n int) int {
	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = -1
	}
	return memoize(n, dp)
}
func memoize(n int, dp []int) int {
	if n <= 1 {
		dp[n] = n
		return dp[n]
	}
	if n == 2 {
		dp[n] = 1
		return dp[n]
	}
	if dp[n] != -1 {
		return dp[n]
	}
	return memoize(n-1, dp) + memoize(n-2, dp) + memoize(n-3, dp)
}

func tabulizeTribonacci(n int) int {
	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	if n <= 1 {
		return n
	}
	if n == 2 {
		return 1
	}
	return tabulize(n, dp)
}
func tabulize(n int, dp []int) int {

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}

func optimizeTribonacci(n int) int {

	var first, second, third int
	if n <= 1 {
		return n
	}
	if n == 2 {
		return 1
	}
	first = 0
	second = 1
	third = 1
	return optimize(n, first, second, third)
}
func optimize(n int, first, second, third int) int {

	for i := 3; i <= n; i++ {
		temp := third
		third = first + second + third
		first = second
		second = temp

	}
	return third
}
