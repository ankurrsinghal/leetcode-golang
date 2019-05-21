package min_stack

import (
	"testing"
)

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func TestMinStack(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)

	if stack.GetMin() != -3 {
		t.Error("GetMin failed -3")
	}

	stack.Pop()

	if stack.Top() != 0 {
		t.Error("Top failed")
	}

	if stack.GetMin() != -2 {
		t.Error("GetMin failed -2")
	}
}
