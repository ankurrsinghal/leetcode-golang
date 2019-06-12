package two_sum_iv_input_is_a_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	return s != nil && (isIdentical(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t))
}

func isIdentical(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	return node1.Val == node2.Val && isIdentical(node1.Left, node2.Left) && isIdentical(node1.Right, node2.Right)
}
