package main

import (
	"errors"
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {

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
	var reversed [][]int
	for i := len(res) - 1; i >= 0; i-- {
		reversed = append(reversed, res[i])
	}
	return reversed
}

type Queue struct {
	node     *TreeNode
	elements []*TreeNode
	size     int
}

// New returns an initialized and empty Queue.
func (q *Queue) New(node *TreeNode) *Queue {
	return &Queue{elements: make([]*TreeNode, 0), node: &TreeNode{}}
}

// Add values to the Queue, as per order of given values.
func (q *Queue) Add(vals ...*TreeNode) {
	q.elements = append(q.elements, vals...)
	q.size += len(vals)
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
