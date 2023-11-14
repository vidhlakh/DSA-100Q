package main

import (
	"fmt"
	"sort"
)

// https://practice.geeksforgeeks.org/problems/power-set4302/1
func main() {
	str := "abc"
	N := len(str)
	//in := strings.Split(str, "")
	var res []string
	possibleSubsequence(str, &res, 0, N, "")
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	// As first element is empty, remove it
	fmt.Printf("Subsequence of %s : %v", str, res[1:])
}

func possibleSubsequence(in string, res *[]string, index, N int, curr string) {
	fmt.Println("index:", index, "curr:", curr)
	if index == N {
		*res = append(*res, curr)
		return
	}

	// include
	possibleSubsequence(in, res, index+1, N, curr+string(in[index]))
	// exclude
	possibleSubsequence(in, res, index+1, N, curr)
}
