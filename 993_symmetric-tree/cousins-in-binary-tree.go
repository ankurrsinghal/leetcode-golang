package cousins_in_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	px, lx := dfs(root, nil, x, 0)
	py, ly := dfs(root, nil, y, 0)
	return px != nil && py != nil && lx == ly && px != py
}

func dfs(node *TreeNode, parent *TreeNode, key int, level int) (*TreeNode, int) {
	if node != nil {
		if node.Val == key {
			return parent, level
		} else {
			l, _ := dfs(node.Left, node, key, level+1)
			r, _ := dfs(node.Right, node, key, level+1)
			if l != nil {
				return l, 0
			} else if r != nil {
				return r, 0
			} else {
				return nil, 0
			}
		}
	}

	return nil, 0
}
