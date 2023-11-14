package main

import (
	"fmt"
	"strings"
)

func main() {

	input := "stets"

	in := strings.Split(input, "")
	N := len(in)

	if palindrome(in, 0, N-1) {
		fmt.Println("String is palindrome")

	} else {
		fmt.Println("String is not a palindrome")

	}
}

func palindrome(in []string, l, h int) bool {

	if l >= h || in[l] != in[h] {
		return false
	}
	palindrome(in, l+1, h-1)
	return true
}
