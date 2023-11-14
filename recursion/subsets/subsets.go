package main

import (
	"fmt"
	"sort"
)

//https://practice.geeksforgeeks.org/problems/subsets-1613027340/1

func main() {
	str := []int{1, 2, 3}
	N := len(str)

	var res []string
	subsets(str, &res, 0, N, "")
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	// As first element is empty, remove it
	fmt.Printf("Subsequence of %v : %v", str, res)
}

func subsets(in []int, res *[]string, index, N int, curr string) {
	fmt.Println("index:", index, "curr:", curr)
	if index == N {
		*res = append(*res, curr)
		return
	}

	// include
	subsets(in, res, index+1, N, fmt.Sprintf("%s %d", curr, in[index]))
	// exclude
	subsets(in, res, index+1, N, curr)
}
