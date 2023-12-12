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
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	levelOrder(tree)
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var curr []int
	var q *Queue[*TreeNode]
	q = q.New(root)
	helper(root, &res, curr, q)
	return res
}

func helper(root *TreeNode, res *[][]int, curr []int, q *Queue[*TreeNode]) {
	if root == nil {
		temp := make([]int, len(curr))
		copy(temp, curr)
		*res = append(*res, temp)
	}
	// Queue is required because on every level, we insert root and while we pop, we insert its childern
	//var queue *TreeNode
	size := q.Size()
	for i := 0; i < size; i++ {
		val, err := q.Dequeue()
		if err != nil {
			fmt.Println("er", err)
		}
		curr = append(curr, val.Val)
		if root.Left != nil {
			helper(root.Left, res, curr, q.Enqueue(root.Val))
		}
		if root.Right != nil {
			helper(root.Right, res, curr, q.Enqueue(root.Val))
		}

	}
}
