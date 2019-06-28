package minimum_distance_between_bst_nodes

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {

	if root == nil {
		return math.MaxInt8
	}

	p := diff(root)
	l := minDiffInBST(root.Left)
	r := minDiffInBST(root.Right)

	if p <= l && p <= r {
		return p
	} else if l <= p && l <= r {
		return l
	}

	return r
}

func diff(node *TreeNode) int {

	if node == nil {
		return math.MaxInt8
	}

	l := node.Val - max(node.Left)
	r := min(node.Right) - node.Val

	if l > r {
		return r
	}

	return l
}

func max(node *TreeNode) int {
	if node == nil {
		return math.MinInt8
	}

	if isLeaf(node) {
		return node.Val
	}

	return max(node.Right)
}

func min(node *TreeNode) int {
	if node == nil {
		return math.MaxInt8
	}

	if isLeaf(node) {
		return node.Val
	}

	return min(node.Left)
}

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}
