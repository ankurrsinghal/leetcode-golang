package cousins_in_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	dx, px := dfs(root, x, 0, nil)
	dy, py := dfs(root, y, 0, nil)

	return dx == dy && px != py
}

func dfs(node *TreeNode, x int, d int, p *TreeNode) (int, *TreeNode) {
	if node == nil {
		return -1, nil
	}

	if node.Val == x {
		return d, p
	}

	dl, pl := dfs(node.Left, x, d+1, node)
	dr, pr := dfs(node.Right, x, d+1, node)

	if dl != -1 {
		return dl, pl
	}

	if dr != -1 {
		return dr, pr
	}

	return -1, nil
}
