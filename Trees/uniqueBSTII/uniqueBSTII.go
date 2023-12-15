package main

// https://leetcode.com/problems/unique-binary-search-trees-ii/
// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	generateTrees(5)
}

func generateTrees(n int) []*TreeNode {

	return helper(1, n)
}
func helper(start, end int) []*TreeNode {
	var res []*TreeNode
	if start > end {
		res = append(res, nil)
	}
	if start == end {
		r := &TreeNode{Val: start}
		res = append(res, r)
		return res
	}
	for i := start; i <= end; i++ {
		// left
		left := helper(start, i-1)
		// right
		right := helper(i+1, end)
		//root := &TreeNode{}
		for _, lnode := range left {
			for _, rnode := range right {
				root := &TreeNode{Val: i}
				root.Left = lnode
				root.Right = rnode
				res = append(res, root)
			}
		}
	}
	return res
}
