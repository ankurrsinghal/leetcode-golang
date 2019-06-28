package path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	return help(root, sum)
}

func help(node *TreeNode, sum int) bool {
	if node == nil {
		return false
	}

	if isLeaf(node) {
		return sum-node.Val == 0
	}

	return help(node.Left, sum-node.Val) || help(node.Right, sum-node.Val)
}

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}
