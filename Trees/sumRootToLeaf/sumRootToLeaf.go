package main

import "fmt"

// https://leetcode.com/problems/sum-root-to-leaf-numbers/
// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println("Sum of values from Root to leaves:", sumNumbers(tree))
}
func sumNumbers(root *TreeNode) int {

	return helper(root, 0)
}

// 1 ,
// 10+2 12 +
// 10+3 13
func helper(root *TreeNode, sum int) int {

	if root == nil {
		return 0
	}

	if root.Left != nil || root.Right != nil {
		return helper(root.Left, sum*10+root.Val) + helper(root.Right, sum*10+root.Val)
	}

	return sum*10 + root.Val // couldnt figure this out

}
