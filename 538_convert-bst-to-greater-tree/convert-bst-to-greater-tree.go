package convert_bst_to_greater_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	helper(root)
	return root
}

func helper(node *TreeNode) {
	if node == nil {
		return
	}

	if node.Right != nil {
		helper(node.Right)
		node.Val = node.Val + node.Right.Val
	}

	if node.Left != nil {
		node.Left.Val = node.Left.Val + node.Val
		helper(node.Left)
	}
}
