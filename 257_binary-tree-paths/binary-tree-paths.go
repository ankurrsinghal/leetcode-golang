package binary_tree_paths

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	a := ""
	r := []string{}
	helper(root, &r, a)
	return r
}

func helper(node *TreeNode, r *[]string, a string) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		*r = append(*r, a+strconv.Itoa(node.Val))
		return
	}
	if node.Left != nil {
		helper(node.Left, r, a+strconv.Itoa(node.Val)+"->")
	}
	if node.Right != nil {
		helper(node.Right, r, a+strconv.Itoa(node.Val)+"->")
	}
}
