package lc

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTreeNaive(root *TreeNode) bool {
	if root.Val == root.Left.Val+root.Right.Val {
		return true
	}
	return false
}

func checkTree(root *TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}
