package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	fmt.Printf("Tree:%+v", buildTree(preorder, inorder))

}
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	mp := make(map[int]int, n)
	for i := 0; i < n; i++ {
		mp[inorder[i]] = i
	}
	return helper(preorder, 0, n-1, inorder, 0, n-1, mp)
}

func helper(preorder []int, pre_st, pre_end int, inorder []int, in_st, in_end int, mp map[int]int) *TreeNode {

	if pre_st > pre_end || in_st > in_end {
		return nil
	}
	root := &TreeNode{
		Val: preorder[pre_st],
	}

	in_root_idx := mp[preorder[pre_st]]
	nums_in_left := in_root_idx - in_st

	// left subtree made mistake in finding pre_end
	root.Left = helper(preorder, pre_st+1, pre_st+nums_in_left, inorder, in_st, in_root_idx-1, mp)

	// right subtree
	root.Right = helper(preorder, pre_st+nums_in_left+1, pre_end, inorder, in_root_idx+1, in_end, mp)

	return root
}
