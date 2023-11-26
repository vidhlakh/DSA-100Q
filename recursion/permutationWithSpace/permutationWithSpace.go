package main

import (
	"fmt"
)

// https://practice.geeksforgeeks.org/problems/permutation-with-spaces3627/1

func main() {
	fmt.Println("Reverse a stringPossible permutations for a string")
	input := "ABC"
	var res []string
	N := len(input)
	permutationWithSpace(input, &res, 0, N)
	// sort.Slice(res, func(i, j int) bool {
	// 	return res[i] < res[j]
	// })
	fmt.Printf("Permutations for %s is %s", input, res)

}

func permutationWithSpace(input string, res *[]string, index, N int, curr string) {

	if index == N {
		*res = append(*res, input)
		return
	}
	//for i := index; i < N; i++ {
	// fmt.Printf("b4-nospace,index:%d,input:%s,res:%v\n", index, input, *res)
	// permutationWithSpace(input, res, index+1, N)
	// fmt.Printf("after-nospace,index:%d,input:%s,res: %v\n", index, input, *res)

	//with space

	fmt.Printf("b4-space,index:%d,input:%s,res:%v\n", index, input, *res)
	permutationWithSpace(input, res, index+1, N, fmt.Sprintf(""))
	fmt.Printf("after-space,index:%d,input:%s,res: %v\n", index, input, *res)
	//}
}
