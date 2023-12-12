package main

import "fmt"

// https://leetcode.com/problems/path-sum-ii/submissions/
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
	targetSum := 22
	fmt.Println("PathSum ", pathSum(tree, targetSum))
}
func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var curr []int
	helper(root, targetSum, &res, curr)

	// if root != nil {helper2(root, targetSum, &res, curr)}
	return res
}

func helper(root *TreeNode, sum int, res *[][]int, curr []int) {
	// HAving this check genericaly makes the code cleaner and easier
	if root == nil {
		return
	}
	curr = append(curr, root.Val)
	if root.Left == nil && root.Right == nil && sum == root.Val {
		temp := make([]int, len(curr))
		copy(temp, curr)
		*res = append(*res, temp)
	}

	helper(root.Left, sum-root.Val, res, curr)
	helper(root.Right, sum-root.Val, res, curr)
}

func helper2(root *TreeNode, sum int, res *[][]int, curr []int) {

	// Instead of checking root is nil at the beginning, we can also check just before we go for left and right
	// This way , we reduce the multiple calls for going to left and right. It wont don that if left or right is nil
	curr = append(curr, root.Val)
	if root.Left == nil && root.Right == nil && sum == root.Val {
		temp := make([]int, len(curr))
		copy(temp, curr)
		*res = append(*res, temp)
	}

	if root.Left != nil {
		helper(root.Left, sum-root.Val, res, curr)
	}
	if root.Left != nil {
		helper(root.Right, sum-root.Val, res, curr)
	}
}
