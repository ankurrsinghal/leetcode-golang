package lowest_common_ancestor_of_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if (root.Left == p || root.Right == p) && (root.Left == q || root.Right == q) {
		return root
	}

	if root.Left == p {
		if search(root.Right, q) {
			return root
		}

		return root.Left
	}

	if root.Right == p {
		if search(root.Left, q) {
			return root
		}

		return root.Right
	}

	if root.Right == q {
		if search(root.Left, p) {
			return root
		}

		return root.Right
	}

	if root.Left == q {
		if search(root.Right, p) {
			return root
		}

		return root.Left
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Left, p, q)

	if l != nil {
		return l
	}

	return r
}

func search(node *TreeNode, q *TreeNode) bool {
	if node == nil {
		return false
	}

	if node == q {
		return true
	}

	return search(node.Left, q) || search(node.Right, q)
}
