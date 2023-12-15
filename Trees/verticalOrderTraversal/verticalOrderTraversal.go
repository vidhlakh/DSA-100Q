package main

import (
	"errors"
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
	// 3 9 20 15 7
	res := verticalTraversal(tree)
	fmt.Println("Vertical Order Traversal:", res)
}

type pair [][]int

type Queue struct {
	node []*TreeNode
	dist []pair
	size int
}

func (q *Queue) Size() int {
	if q == nil {
		return 0
	}
	// length of node and dist is expected to be same
	if len(q.node) != len(q.dist) {
		return -1
	}
	return len(q.node)
}
func (q *Queue) Enqueue(node *TreeNode, dist pair) error {
	if q == nil {
		return errors.New("queue is not initialized")
	}
	q.node = append(q.node, node)
	q.dist = append(q.dist, dist)
	q.size = q.Size() + 1
	return nil
}
func (q *Queue) Dequeue() (*TreeNode, pair) {
	var (
		node *TreeNode
		dist pair
	)
	node = q.node[len(q.node)-1]
	dist = q.dist[len(q.dist)-1]

	q.node = q.node[:len(q.node)-1]
	q.dist = q.dist[:len(q.dist)-1]
	q.size = q.Size() - 1
	return node, dist
}
func verticalTraversal(root *TreeNode) [][]int {
	var res [][]int

	var q *Queue
	for q.Size() != 0 {
		s := q.Size()
		for i := 0; i < s; i++ {
			node, dist := q.Dequeue()

			if node.Left != nil {
				q.Enqueue(node.Left, dist)
			}
			if node.Right != nil {
				q.Enqueue(node.Right, dist)
			}
			curr := make([][]int, 0)
			fmt.Println("curr", curr)
			curr = append(curr, dist...)
			res = append(res, curr...)
		}
	}
	fmt.Printf("Result: %+v", res)
	return res
}
