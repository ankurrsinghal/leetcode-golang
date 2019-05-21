package backspace_string_compare

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
func TestBackspaceCompareSuccess(t *testing.T) {
	result := backspaceCompare("ak#c", "ad#c")

	if !result {
		t.Error("Failed")
	}
}

func TestBackspaceCompareFail(t *testing.T) {
	result := backspaceCompare("bk#c", "ad#c")

	if result {
		t.Error("Failed")
	}
}
