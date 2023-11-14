package main

import (
	"fmt"
)

func main() {
	fmt.Println("Reverse a stringPossible permutations for a string")
	input := "ABC"
	var res []string
	N := len(input)
	permutation(input, &res, 0, N)

	fmt.Printf("Permutations for %s is %s", input, res)

}

func permutation(input string, res *[]string, index, N int) {

	if index == N-1 {
		*res = append(*res, input)
		return
	}
	for i := index; i < N; i++ {
		fmt.Printf("b4,index:%d,i:%d,input:%s,res:%v\n", index, i, input, *res)
		inRune := []rune(input)
		inRune[index], inRune[i] = inRune[i], inRune[index]
		permutation(string(inRune), res, index+1, N)
		// no swap required again as we didnt modify input directly
		// inRune[index], inRune[i] = inRune[i], inRune[index]
		// input = string(inRune)
		fmt.Printf("after,index:%d,i:%d,input:%s,res: %v\n", index, i, input, *res)
	}
}
