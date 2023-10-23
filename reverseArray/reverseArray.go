package main

import "fmt"

func main() {
	fmt.Println("REVERSE Array")
	reverseArrayMethod1()
	reverseArrayMethod2()
}

// Space complexity :O(1)   Time complexity :O(n)
func reverseArrayMethod1() {
	input := []string{"10", "20", "30", "40", "50", "60"}
	start := 0
	end := len(input) - 1
	fmt.Printf("Address: %p\n", &input)
	for start < end {
		temp := input[start]
		input[start] = input[end]
		input[end] = temp
		start++
		end--
	}
	fmt.Printf("Address after reverse : %p\n", &input)
	fmt.Println("reversed:", input)
}

// Space complexity :O(n)   Time complexity :O(n)
func reverseArrayMethod2() {
	input := []string{"10", "20", "30", "40", "50", "60"}
	start := 0
	end := len(input) - 1
	fmt.Printf("Address : %p\n", &input)
	recursiveReverseArray(input, start, end)
	fmt.Println("reversedmethod2:", input)
}

func recursiveReverseArray(input []string, start, end int) {
	fmt.Printf("Address in recurse : %d,%p\n", start, &input)
	if start > end {
		return
	}
	temp := input[start]
	input[start] = input[end]
	input[end] = temp
	recursiveReverseArray(input, start+1, end-1)
}
