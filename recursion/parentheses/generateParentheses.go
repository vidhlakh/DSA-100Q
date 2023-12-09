package main

import "fmt"

// https://leetcode.com/problems/generate-parentheses/

func main() {
	// Method1
	fmt.Println("Generate wellformed parentheses:", generateParenthesesM1(3))
}

func generateParenthesesM1(n int) []string {

	var curr string
	var res []string
	helperM1(n, 0, 0, &res, curr)
	return res
}

//  0 to n approach
func helperM1(n, open, closed int, res *[]string, curr string) {
	if closed == n {
		*res = append(*res, curr)
		return
	}
	// add open paren
	if open < n {
		helperM1(n, open+1, closed, res, curr+"(")
	}
	// add closed paren
	if open > closed {
		helperM1(n, open, closed+1, res, curr+")")
	}
}
