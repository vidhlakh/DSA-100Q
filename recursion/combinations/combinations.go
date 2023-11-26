package main

import (
	"fmt"
)

// https://leetcode.com/problems/combinations/description/

func main() {
	n := 4
	k := 2
	//fmt.Printf("Method 1 Combinations for n = %d are %v\n", n, combineMethod1(n, k))

	fmt.Printf("Method 2 Combinations for n = %d are %v\n", n, combineMethod2(n, k))
}

// Method 1 n to 0
// func combineMethod1(n int, k int) [][]int {
// 	var res [][]int
// 	var curr []int
// 	for i := n; i > 1; i-- {

// 		//curr = append(curr, i)
// 		helperMethod1(i, k, &res, curr)
// 	}

// 	return res
// }

// func helperMethod1(ind int, k int, res *[][]int, curr []int) {

// 	if ind <= 0 {
// 		return
// 	}

// 	// include
// 	curr = append(curr, ind)
// 	helperMethod1(ind-1, k, res, curr)
// 	*res = append(*res, curr)

// 	// // exclude
// 	// helperMethod1(ind-1,k,res,curr)
// }

// Method 2 1 to n
func combineMethod2(n, k int) [][]int {
	var res [][]int
	var curr []int
	helperMethod2(1, n, k, &res, curr)
	fmt.Println("res", res)
	return res
}

func helperMethod2(i, n, k int, res *[][]int, curr []int) {

	if len(curr) == k {

		*res = append(*res, curr)
		fmt.Println("res", *res)
		return
	}
	if i > n {
		return
	}

	// pick
	curr = append(curr, i)
	helperMethod2(i+1, n, k, res, curr)
	curr = curr[:len(curr)-1]

	// dont pick
	helperMethod2(i+1, n, k, res, curr)

}
