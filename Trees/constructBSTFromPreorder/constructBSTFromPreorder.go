package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	preorder := []int{8, 5, 1, 7, 10, 12}
	fmt.Printf("BST From Preorder: %+v", bstFromPreorder(preorder))
}
func bstFromPreorder(preorder []int) *TreeNode {
	return helper(preorder, 0, len(preorder)-1)
}
func helper(preorder []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	// first is root
	root := &TreeNode{
		Val: preorder[start],
	}
	idx := start
	for idx = start; idx <= end; idx++ {
		if preorder[idx] > preorder[start] {
			break
		}
	}

	// Go over the preorder arr and elements belong to left until the element is bigger than the root
	root.Left = helper(preorder, start+1, idx-1)
	root.Right = helper(preorder, idx, end)
	fmt.Printf("rr:%+v", root)
	return root
}
