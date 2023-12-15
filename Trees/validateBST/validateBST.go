package main

import (
	"fmt"
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
		},
		Right: &TreeNode{
			Val: 13,
			Left: &TreeNode{
				Val: 12,
			},
			Right: &TreeNode{
				Val: 17,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 19,
				},
			},
		},
	}
	fmt.Println("Validate BST ", isValidBST(tree))
}
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, min, max int) bool {

	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}
	// go through the validation
	return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)

}
