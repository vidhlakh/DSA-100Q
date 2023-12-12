package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var res1, res2, res3 []int
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
	preorderTraversal(tree, &res1)
	fmt.Println("Preorder Traversal", res1)

	inorderTraversal(tree, &res2)
	fmt.Println("Inorder Traversal", res2)

	postorderTraversal(tree, &res3)
	fmt.Println("Postorder Traversal", res3)

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
