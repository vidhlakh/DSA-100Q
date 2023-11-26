package main

import "fmt"

// https://leetcode.com/problems/decode-ways/description/

func main() {

	// Method 1
	fmt.Println("Decode ways method1", numDecodingsM1("12"))

	// Method 2
	fmt.Println("Decode ways method2", numDecodingsM2("12"))

}

func numDecodingsM1(s string) int {
	if string(s[0]) == "0" {
		return 0
	}
	return helperM1(len(s)-1, s)
}

// n to 0
func helperM1(ind int, s string) int {

	if ind <= 0 {
		return 1
	}
	var res int
	// single digit
	if string(s[ind]) != "0" {
		res = helperM1(ind-1, s)
	}
	// double digit
	if string(s[ind-1]) == "1" || (string(s[ind-1]) == "2" && string(s[ind]) >= "0" && string(s[ind]) <= "6") {
		res += helperM1(ind-2, s)
	}
	return res
}

func numDecodingsM2(s string) int {
	if string(s[0]) == "0" {
		return 0
	}
	return helperM2(0, s, len(s))
}

// 0 to n
func helperM2(ind int, s string, n int) int {

	if ind == n {
		return 1
	}
	var res int
	// single digit
	if string(s[ind]) != "0" {
		res = helperM2(ind+1, s, n)
	}
	// double digit
	if (ind < n-1) && (string(s[ind]) == "1" || (string(s[ind]) == "2" && string(s[ind+1]) >= "0" && string(s[ind+1]) <= "6")) {
		res += helperM2(ind+2, s, n)
	}
	return res
}
