package main

import "fmt"

func main() {
	//var n *int
	N := 16
	var res1, res2 []int
	// Method 1
	printPatternMethod1(&res1, N, N, false)
	fmt.Println("res1", res1)

	//Method2
	printPatternMethod2(&res2, N)
	fmt.Println("res2", res2)

}

// With flag and N
func printPatternMethod1(res *[]int, N, present int, incr bool) {
	*res = append(*res, present)
	if incr {
		if present == N {
			return
		} else {
			printPatternMethod1(res, N, present+5, incr)
		}

	} else {
		if present-5 > 0 {
			printPatternMethod1(res, N, present-5, false)
		} else {
			printPatternMethod1(res, N, present-5, true)
		}

	}

}

// Without flag and N
func printPatternMethod2(res *[]int, present int) {
	//temp := *res

	if present < 0 {
		*res = append(*res, present)
		return
	}

	*res = append(*res, present)
	printPatternMethod2(res, present-5)
	*res = append(*res, present)

}
