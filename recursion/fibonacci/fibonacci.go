package main

import "fmt"

func main() {
	n := 7
	res := fib(n)
	fmt.Printf("Recursion::%dth element of Fibonacci series is %d\n", n, res)

	// Memoization
	res = fibMemoize(n)
	fmt.Printf("Memoization:: %dth element of Fibonacci series is %d\n", n, res)

	// Tabulation
	res = fibTabulation(n)
	fmt.Printf("Tabulation:: %dth element of Fibonacci series is %d\n", n, res)

	// Space optimization
	res = fibSpaceOptimization(n)
	fmt.Printf("Space Optimization:: %dth element of Fibonacci series is %d\n", n, res)

}

// recursion
func fib(i int) int {
	//fmt.Printf("%d\t", i)
	if i <= 1 {
		return i
	}
	return fib(i-1) + fib(i-2)

}

// Memoization (Top down approach)
func fibMemoize(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -1
	}
	memoize(n, dp)
	return dp[n]
}

func memoize(i int, dp []int) int {

	if i <= 1 {
		return i
	}
	if dp[i] != -1 {
		return dp[i]
	}
	dp[i] = memoize(i-1, dp) + memoize(i-2, dp)
	return dp[i]
}

// Tabulation (Bottom up approach)
func fibTabulation(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1
	return tabulize(n, dp)

}

func tabulize(n int, dp []int) int {

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]

}

// Space Optimization
func fibSpaceOptimization(n int) int {
	if n <= 1 {
		return n
	}
	prev := 0
	curr := 1
	return optimize(n, prev, curr)

}

func optimize(n, prev, curr int) int {

	for i := 2; i <= n; i++ {
		temp := curr
		curr = prev + curr
		prev = temp
	}
	return curr

}
