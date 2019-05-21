package backspace_string_compare

import "strings"

// Stack ...
type StringStack struct {
	data string
}

func (stack *StringStack) isEmpty() bool {
	return len(stack.data) == 0
}

func (stack *StringStack) size() int {
	return len(stack.data)
}

func (stack *StringStack) pop() {
	if stack.size() == 0 {
		return
	}
	stack.data = stack.data[:stack.size()-1]
}

func (stack *StringStack) top() string {
	return string(stack.data[stack.size()-1])
}

func (stack *StringStack) push(str string) {
	stack.data += str
}

func createStringStack() *StringStack {
	return &StringStack{data: ""}
}

func backspaceCompare(S string, T string) bool {

	stack1 := createStringStack()
	for _, s := range S {
		str := string(s)
		if str == "#" {
			stack1.pop()
		} else {
			stack1.push(str)
		}
	}

	stack2 := createStringStack()
	for _, s := range T {
		str := string(s)
		if str == "#" {
			stack2.pop()
		} else {
			stack2.push(str)
		}
	}

	return strings.Compare(stack1.data, stack2.data) == 0
}
