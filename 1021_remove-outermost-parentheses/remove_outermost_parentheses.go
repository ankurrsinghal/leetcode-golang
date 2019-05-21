package remove_outermost_parentheses

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
	if s.size() == s.cap {
		if s.size() == 0 {
			s.data = make([]string, 1)
			s.cap = 1
		} else {
			s.data = append(s.data, make([]string, s.cap*2)...)
			s.cap *= 2
		}
	}

	s.pointer++
	s.data[s.pointer] = str
}

func createStack(cap int) *Stack {
	return &Stack{cap: cap, pointer: -1}
}

func removeOuterParentheses(S string) string {

	if S == "" {
		return S
	}

	stack := createStack(0)
	var result []string

	lastI := 0
	for i, paren := range S {
		char := string(paren)
		if stack.isEmpty() {
			if i != 0 {
				result = append(result, S[lastI+1:i-1])
			}
			stack.push(char)
			lastI = i
		} else if stack.top() == "(" && ")" == char {
			stack.pop()
		} else {
			stack.push(char)
		}
	}

	result = append(result, S[lastI+1:len(S)-1])

	return strings.Join(result, "")
}
