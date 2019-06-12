package univalued_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	var result int
	dfs(root, L, R, &result)
	return result
}

func dfs(node *TreeNode, L int, R int, result *int) {

	if node == nil {
		return
	}

	if L <= node.Val && R >= node.Val {
		*result += node.Val
	}

	if L < node.Val {
		dfs(node.Left, L, R, result)
	}

	if R > node.Val {
		dfs(node.Right, L, R, result)
	}
}
