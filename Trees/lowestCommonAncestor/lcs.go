package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

func lowestCommonAncestorBT(root, p, q *TreeNode) *TreeNode {

	if root == p || root == q || root == nil {
		return root
	}
	//fmt.Println("root.Val",root.Val)
	right := lowestCommonAncestorBT(root.Right, p, q)
	left := lowestCommonAncestorBT(root.Left, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
