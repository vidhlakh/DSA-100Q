package main

import (
	"fmt"
	"sort"
)

//https://practice.geeksforgeeks.org/problems/subsets-1613027340/1
// https://leetcode.com/problems/subsets/

func main() {
	str := []int{1, 2, 3}
	N := len(str)

	var res []string
	subsetsMethod1(str, &res, 0, N, "")
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	// As first element is empty, remove it
	fmt.Printf("Subsequence of %v : %v", str, res)

	// Subset with response [][]int
	// https://leetcode.com/problems/subsets/
	fmt.Printf("Subset of input: %v is %v\n", str, subsets(str))

}

func subsetsMethod1(in []int, res *[]string, index, N int, curr string) {
	fmt.Println("index:", index, "curr:", curr)
	if index == N {
		*res = append(*res, curr)
		return
	}

	// include
	subsetsMethod1(in, res, index+1, N, fmt.Sprintf("%s %d", curr, in[index]))
	// exclude
	subsetsMethod1(in, res, index+1, N, curr)
}

// Discussed on Day 2 Recursion class

func subsets(nums []int) [][]int {
	var res [][]int
	var cur []int
	helper(nums, 0, &res, cur)
	return res
}

func helper(nums []int, ind int, res *[][]int, cur []int) {
	if ind == len(nums) {
		*res = append(*res, cur)
		return
	}
	// pick
	cur = append(cur, nums[ind])
	helper(nums, ind+1, res, cur)

	// dont pick
	cur = cur[:len(cur)-1]
	helper(nums, ind+1, res, cur)

}
