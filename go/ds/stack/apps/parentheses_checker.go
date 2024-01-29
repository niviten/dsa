package apps

import "dsa/ds/stack"

func IsParenthesesBalanced(input string) bool {
	s := stack.NewStack[rune]()
	for _, ch := range input {
		if isOpenParentheses(ch) {
			s.Push(ch)
		} else if isCloseParentheses(ch) {
			popped, err := s.Pop()
			if err != nil {
				return false
			}
			if !isMatch(popped, ch) {
				return false
			}
		}
	}
	return s.IsEmpty()
}

func isOpenParentheses(ch rune) bool {
	return ch == '(' || ch == '[' || ch == '{'
}

func isCloseParentheses(ch rune) bool {
	return ch == ')' || ch == ']' || ch == '}'
}

func isMatch(open, close rune) bool {
	if open == '(' && close == ')' {
		return true
	}
	if open == '[' && close == ']' {
		return true
	}
	if open == '{' && close == '}' {
		return true
	}
	return false
}
