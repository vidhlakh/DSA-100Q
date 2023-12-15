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
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Printf("Insert into BST: %+v", insertIntoBST(tree, 10))
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	return helper(root, val, math.MinInt64, math.MaxInt64)
}
func helper(root *TreeNode, val int, min, max int) *TreeNode {

	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = helper(root.Left, val, min, val)
	} else {
		root.Right = helper(root.Right, val, val, max)
	}
	return root
}
