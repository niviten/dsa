package apps

import (
	"testing"
)

var paranthesesCheckerTestCases = []struct {
	input string
	want  bool
}{
	{"", true},
	{"()", true},
	{"([])", true},
	{"{[]}", true},
	{"(}", false},
	{"([)]", false},
	{"((()", false},
	{"asdfv", true},
}

func TestIsParenthesesBalanced(t *testing.T) {
	for _, testCase := range paranthesesCheckerTestCases {
		got := IsParenthesesBalanced(testCase.input)
		if got != testCase.want {
			t.Errorf("IsParenthesesBalanced(%q) = %v; want %v", testCase.input,
				got, testCase.want)
		}
	}
}
