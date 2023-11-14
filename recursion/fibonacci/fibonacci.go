package main

import "fmt"

func main() {
	n := 5
	res := fib(n)
	fmt.Printf("%dth element of Fibonacci series is %d\n", n, res)
}
func fib(i int) int {
	fmt.Printf("%d\t", i)
	if i <= 1 {
		return i
	}
	return fib(i-1) + fib(i-2)

}
