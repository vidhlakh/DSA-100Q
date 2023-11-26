package main

import "fmt"

// https://leetcode.com/problems/climbing-stairs
func main() {
	fmt.Println(climbStairs(3))
}
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}
