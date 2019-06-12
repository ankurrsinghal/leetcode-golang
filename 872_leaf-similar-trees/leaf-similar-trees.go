package convert_bst_to_greater_tree

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	p := ""
	q := ""

	inOrder(root1, &p)
	inOrder(root2, &q)

	return p == q
}

func inOrder(node *TreeNode, s *string) {
	if node == nil {
		return
	}

	inOrder(node.Left, s)

	if isLeaf(node) {
		*s += strconv.Itoa(node.Val)
	}

	inOrder(node.Right, s)
}

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}
