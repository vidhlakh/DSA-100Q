package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Reverse a string")
	input := "Geeks"

	in := strings.Split(input, "")
	N := len(in)
	reverseString(&in, 0, N-1)
	fmt.Printf("Reveresed string for %s is %s", input, strings.Join(in, ""))
}

func reverseString(in *[]string, l, h int) {
	res := *in
	if l >= h {
		return
	}
	res[l], res[h] = res[h], res[l]
	// temp := res[l]
	// res[l] = res[h]
	// res[h] = temp
	reverseString(in, l+1, h-1)
}
