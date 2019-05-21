package remove_all_adjacent_duplicates_in_string

import (
	"strings"
)

// Stack ...
type Stack struct {
	data    []string
	cap     int
	pointer int
}

func (s *Stack) isEmpty() bool {
	return s.pointer == -1
}

func (s *Stack) size() int {
	return s.pointer + 1
}

func (s *Stack) pop() {
	s.data[s.pointer] = ""
	s.pointer--
}

func (s *Stack) top() string {
	return s.data[s.pointer]
}

func (s *Stack) push(str string) {
	s.pointer++
	s.data[s.pointer] = str
}

func createStack(cap int) *Stack {
	return &Stack{data: make([]string, cap), cap: cap, pointer: -1}
}

func removeDuplicates(S string) string {

	if S == "" {
		return S
	}

	stack := createStack(20000)
	for _, character := range S {
		char := string(character)

		if !stack.isEmpty() && stack.top() == char {
			stack.pop()
		} else {
			stack.push(char)
		}
	}

	return strings.Join(stack.data, "")
}
