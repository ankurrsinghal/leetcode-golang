package univalued_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	return dfs(root, root.Val)
}

func dfs(node *TreeNode, val int) bool {
	if node == nil {
		return true
	}

	if node.Val == val {
		return dfs(node.Left, val) && dfs(node.Right, val)
	}

	return false
}
