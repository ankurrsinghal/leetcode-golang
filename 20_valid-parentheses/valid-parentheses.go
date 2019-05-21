package valid_parentheses

const tightConstant = 10

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
			s.data = append(s.data, make([]string, s.cap+tightConstant)...)
			s.cap += tightConstant
		}
	}

	s.pointer++
	s.data[s.pointer] = str
}

func createStack() *Stack {
	return &Stack{cap: 0, pointer: -1}
}

func isValid(s string) bool {
	if s == "" {
		return true
	}

	stack := createStack()
	for _, char := range s {
		if stack.isEmpty() {
			stack.push(string(char))
			continue
		}

		if stack.top() == "(" {
			if string(char) == ")" {
				stack.pop()
				continue
			}
		}

		if stack.top() == "{" {
			if string(char) == "}" {
				stack.pop()
				continue
			}
		}

		if stack.top() == "[" {
			if string(char) == "]" {
				stack.pop()
				continue
			}
		}

		stack.push(string(char))
	}

	return stack.isEmpty()
}
