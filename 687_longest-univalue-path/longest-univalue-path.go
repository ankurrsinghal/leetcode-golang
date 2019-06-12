package longest_univalue_path

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	var count int
	helper(root, "", 0, &count)

	return count
}

func helper(node *TreeNode, path string, lastCount int, maxCount *int) {

	if node == nil {
		return
	}

	str := strconv.Itoa(node.Val)
	if str == path {
		count := lastCount + 1
		if *maxCount < count {
			*maxCount = count
		}
		helper(node.Left, path, count, maxCount)
		helper(node.Right, path, count, maxCount)
		return
	}

	helper(node.Left, str, 0, maxCount)
	helper(node.Right, str, 0, maxCount)
}
