package diameter_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	var tilt int
	sumOfTree(root, &tilt)
	return tilt
}

func sumOfTree(node *TreeNode, tilt *int) int {
	if node == nil {
		return 0
	}

	l := sumOfTree(node.Left, tilt)
	r := sumOfTree(node.Right, tilt)

	*tilt += abs(l - r)

	return node.Val + l + r
}

func abs(val int) int {
	if val > 0 {
		return val
	} else {
		return -1 * val
	}
}
