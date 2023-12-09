package main

import "fmt"

// https://leetcode.com/problems/min-cost-climbing-stairs
func main() {
	fmt.Println("Recursive min cost")
	cost := []int{10, 15, 20}
	res := minCostClimbingStairs(cost)
	fmt.Printf("Min cost to climb stairs: %v, %v\n", cost, res)

	// Method 2
	res2 := minCostClimbingStairsMethod2(cost)
	fmt.Printf("Min cost to climb stairs2: %v, %v\n", cost, res2)
	fmt.Printf("Min cost tabulize:%d", minCostClimbingStairsTabulize(cost))
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)

	return minimum(helperFunc(cost, n-1), helperFunc(cost, n-2))
}

func helperFunc(cost []int, index int) int {
	if index <= 1 {
		return cost[index]
	}

	return cost[index] + minimum(helperFunc(cost, index-1), helperFunc(cost, index-2))
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Method 2

func minCostClimbingStairsMethod2(cost []int) int {

	ind := 0
	return minimum(helperFuncMethod2(cost, ind+1), helperFuncMethod2(cost, ind+2))
}

func helperFuncMethod2(cost []int, index int) int {
	if index >= len(cost) {
		return 0
	}
	if index == len(cost)-1 {
		return cost[index]
	}

	return cost[index] + minimum(helperFuncMethod2(cost, index+1), helperFuncMethod2(cost, index+2))
}

func minCostClimbingStairsTabulize(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0] = cost[0]
	dp[1] = cost[1]
	return tabulize(n, cost, dp)
}
func tabulize(n int, cost, dp []int) int {
	// initial condn

	for i := 2; i < n; i++ {
		dp[i] = cost[i] + minimum(dp[i-1], dp[i-2])
	}
	return minimum(dp[n-1], dp[n-2])

}
