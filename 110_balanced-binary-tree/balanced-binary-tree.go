package balanced_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}

	l := depth(root.Left)
	r := depth(root.Right)

	if r-l > 1 || l-r > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	l := depth(node.Left)
	r := depth(node.Right)

	if l > r {
		return l + 1
	}
	return r + 1
}
