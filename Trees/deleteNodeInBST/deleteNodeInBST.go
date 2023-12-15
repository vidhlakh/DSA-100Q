package main

import (
	"fmt"
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

	res := deleteNode(tree, 15)
	fmt.Println("Delete node in BST ")
	
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val < key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val > key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			temp := root.Right
			root = temp
		} else if root.Right == nil {
			temp := root.Left
			root = temp
		} else {
			// 2 options
			// use successor - min of right subtree
			// predecessor - max of left subtree

			successor := root.Right
			for successor.Left != nil {
				successor = successor.Left
			}
			root.Val = successor.Val
			root.Right = deleteNode(root.Right, successor.Val)
		}
	}
	return root
}
