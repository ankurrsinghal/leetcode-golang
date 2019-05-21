package remove_outermost_parentheses

import (
	"testing"
)

func TestRemoveOuterParentheses(t *testing.T) {
	r := removeOuterParentheses("")

	if r != "" {
		t.Error("Not Equal")
	}
}

func TestRemoveOuterParentheses2(t *testing.T) {
	r := removeOuterParentheses("()()")

	if r != "" {
		t.Error("Not Equal")
	}
}

func TestRemoveOuterParentheses3(t *testing.T) {
	r := removeOuterParentheses("(()())(())(()(()))")

	if r != "()()()()(())" {
		t.Error("Not Equal")
	}
}
