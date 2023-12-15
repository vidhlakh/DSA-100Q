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
	//var res1, res2, res3  []int
	var res4 []int
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
	// preorderTraversal(tree, &res1)
	// fmt.Println("Preorder Traversal", res1)

	// inorderTraversal(tree, &res2)
	// fmt.Println("Inorder Traversal", res2)

	// postorderTraversal(tree, &res3)
	// fmt.Println("Postorder Traversal", res3)

	res4 = preorderTraversalIterative(tree)
	fmt.Println("Preorder Traversal Iterative", res4)

}

// Left Root Right
func inorderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left, res)
	*res = append(*res, root.Val)
	inorderTraversal(root.Right, res)

}

// Root Left Right
func preorderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preorderTraversal(root.Left, res)
	preorderTraversal(root.Right, res)

}

// Left Right Root
func postorderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	postorderTraversal(root.Left, res)
	postorderTraversal(root.Right, res)
	*res = append(*res, root.Val)

}

// Iterative approach

type Stack struct {
	node []*TreeNode
	size int
}

func (s *Stack) Size() int {
	if s == nil {
		return 0
	}
	return len(s.node)
}
func (s *Stack) push(node *TreeNode) error {
	if s == nil {
		return errors.New("stack is nil")
	}
	s.node = append(s.node, node)
	s.size = s.Size() + 1
	return nil
}

func (s *Stack) pop() (*TreeNode, error) {
	if s == nil || len(s.node) < 1 {
		return nil, errors.New("no elements in stack to pop")
	}
	node := s.node[0]
	s.node = s.node[1:]
	s.size = s.Size() - 1
	return node, nil
}
func preorderTraversalIterative(root *TreeNode) []int {
	var res []int
	s := &Stack{}
	err := s.push(root)
	if err != nil {
		fmt.Println("err to push", err)
	}
	for s.Size() != 0 {
		node, err := s.pop()
		if err != nil {
			fmt.Println("err poping", err)
		}
		res = append(res, node.Val)
		if root.Left != nil {
			err := s.push(root.Left)
			if err != nil {
				fmt.Println("err pushing ", err)
			}
		}
		if root.Right != nil {
			err := s.push(root.Right)
			if err != nil {
				fmt.Println("err pushing ", err)
			}
		}
	}
	return res
}
