package main

import (
	"fmt"
	"math"
)

type Node struct {
	Root int
}
type Tree struct {
}

func main() {
	fmt.Println("Kth smallest")
	kthSmallestMethod1()
}

// 1, 3 ,6,12,21,32,34,56,87,343
// Linear sorting can be done through heap sort
// sorting in as pre order traversal
func kthSmallestMethod1() {
	input := []int{1, 3, 343, 21, 12, 56, 6, 34, 87, 32}
	k := 5
	min := int(math.Inf(-1))
	count := 0

	for _, inp := range input {
		fmt.Println("min,inp", min, inp)
		if min > inp {
			min = inp
			count++
		}
		if count == k {
			break
		}

	}
	fmt.Println("Kth smallest", min)
}

func (node *Node) insert(element int) {

	if element < node.Root {
		insert(left, element)

	} else {
		Right(right, element)
	}
}
