package search_in_a_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	i := root
	for i != nil {
		if i.Val == val {
			return i
		}

		if i.Val > val {
			i = i.Left
		} else {
			i = i.Right
		}
	}

	return i
}
