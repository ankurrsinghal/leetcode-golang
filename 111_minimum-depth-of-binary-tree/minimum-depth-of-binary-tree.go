package minimum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)

	if l == 0 {
		return r + 1
	} else if r == 0 {
		return l + 1
	} else if l < r {
		return l + 1
	} else {
		return r + 1
	}
}
