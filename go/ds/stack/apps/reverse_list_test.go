package apps

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	name     string
	input    []int
	expected []int
}{
	{
		name:     "odd size list",
		input:    []int{1, 2, 3, 4, 5},
		expected: []int{5, 4, 3, 2, 1},
	},
	{
		name:     "even size list",
		input:    []int{1, 2, 3, 4, 5, 6},
		expected: []int{6, 5, 4, 3, 2, 1},
	},
	{
		name:     "empty list",
		input:    []int{},
		expected: []int{},
	},
	{
		name:     "single element list",
		input:    []int{5},
		expected: []int{5},
	},
}

func TestReverseList(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ReverseList(testCase.input)
			if !reflect.DeepEqual(testCase.input, testCase.expected) {
				t.Errorf("ReverseList(%v) == %v, want %v",
					testCase.input, testCase.input, testCase.expected)
			}
		})
	}
}
