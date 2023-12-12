package main

import (
	"errors"
	"fmt"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Queue struct {
	elements []*TreeNode
	size     int
}

// New returns an initialized and empty Queue.
func (q *Queue) New(node *TreeNode) *Queue {
	return &Queue{elements: make([]*TreeNode, 0)}
}

// Enqueue adds given value to end of Queue.
// Does nothing if `Queue` is not initialized
func (q *Queue) Enqueue(val *TreeNode) {
	if q == nil {
		return
	}
	q.elements = append(q.elements, val)
	q.size += 1
}

// Dequeue deletes and returns first element from Queue.
// Returns -1 if Queue is un-initialized or has no more elements to remove(empty Queue).
func (q *Queue) Dequeue() (*TreeNode, error) {
	var ret *TreeNode
	if q == nil {
		return ret, errors.New("queue is not initialized or nil")
	}
	if q.Size() == 0 || q.elements == nil {
		return ret, errors.New("queue is not initialized or nil")
	}
	ret = q.elements[0]
	q.elements = q.elements[1:]
	q.size -= 1
	return ret, nil
}

// Peek returns the first element in the Queue.
func (q *Queue) Peek() (*TreeNode, error) {
	if q == nil {
		var ret *TreeNode
		return ret, errors.New("queue is not initialized or nil")
	}
	return q.elements[0], nil
}

// Size returns current size of Queue.
func (q *Queue) Size() int {
	return q.size
}

// Implements stringer interface for Queue
func (q *Queue) String() string {
	var sArr []string
	for _, v := range q.elements {
		sArr = append(sArr, fmt.Sprintf("%v", v))
	}
	return strings.Join(sArr, ",")
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
	res := levelOrder(tree)
	fmt.Println("Level Order Traversal:", res)
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	var q *Queue
	q = q.New(root)

	// Queue is required because on every level, we insert root and while we pop, we insert its childern
	//var queue *TreeNode
	q.Enqueue(root)
	for q.Size() != 0 {
		size := q.Size()
		curr := make([]int, 0)
		for i := 0; i < size; i++ {
			val, err := q.Dequeue()
			if err != nil {
				fmt.Println("er", err)
			}
			curr = append(curr, val.Val)
			if val.Left != nil {
				q.Enqueue(val.Left)

			}
			if val.Right != nil {
				q.Enqueue(val.Right)

			}
		}
		res = append(res, curr)
	}
	return res
}
