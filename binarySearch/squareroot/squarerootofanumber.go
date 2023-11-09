package main

import "fmt"

func main() {
	input := 4
	res := squareroot(input)
	fmt.Printf("Squareroot of %v is %d", input, res)
}

func squareroot(input int) int {
	var l, h, res int
	l = 0
	h = input / 2

	for l <= h {
		mid := l + (h - l) //2
		if mid*mid == input {
			return mid
		} else {
			if mid*mid < input {
				l = mid + 1
				res = mid
			} else {
				h = mid - 1
				res = mid
			}
		}
	}
	return res
}
