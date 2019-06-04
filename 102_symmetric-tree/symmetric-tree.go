package symmetric_tree

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

func isSymmetric(root *TreeNode) bool {
	fmt.Println(inOrder(root))
	fmt.Println(inOrderMirror(root))
	return strings.Compare(inOrder(root), inOrderMirror(root)) == 0
}

func inOrder(node *TreeNode) string {
	if node == nil {
		return ""
	}

	s := ""
	l := inOrder(node.Left)
	r := inOrder(node.Right)

	if l != "" {
		s += fmt.Sprintf("L(%s)", l)
	}

	s += fmt.Sprintf("P(%s)", strconv.Itoa(node.Val))

	if r != "" {
		s += fmt.Sprintf("R(%s)", r)
	}

	return s
}

func inOrderMirror(node *TreeNode) string {
	if node == nil {
		return ""
	}

	s := ""
	l := inOrderMirror(node.Right)
	r := inOrderMirror(node.Left)

	if l != "" {
		s += fmt.Sprintf("L(%s)", l)
	}

	s += fmt.Sprintf("P(%s)", strconv.Itoa(node.Val))

	if r != "" {
		s += fmt.Sprintf("R(%s)", r)
	}

	return s
}
