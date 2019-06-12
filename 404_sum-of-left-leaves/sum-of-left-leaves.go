package sum_of_root_to_leaf_binary_numbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, prefix int) int {
	if node == nil {
		return 0
	}

	prefix = node.Val + prefix<<1

	if node.Left != nil || node.Right != nil {
		return dfs(node.Left, prefix) + dfs(node.Right, prefix)
	}

	return prefix
}
