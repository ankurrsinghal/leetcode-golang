package same_tree

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return strings.Compare(inOrder(p), inOrder(q)) == 0
}

func inOrder(node *TreeNode) string {
	if node == nil {
		return ""
	}

	return fmt.Sprintf("L(%s)", inOrder(node.Left)) + fmt.Sprintf("P(%s)", strconv.Itoa(node.Val)) + fmt.Sprintf("R(%s)", inOrder(node.Right))
}
