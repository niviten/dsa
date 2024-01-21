package sort

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
		name:     "random list",
		input:    []int{3, 1, 4, 1, 5, 9, 2, 6},
		expected: []int{1, 1, 2, 3, 4, 5, 6, 9},
	},
	{
		name:     "sorted list",
		input:    []int{1, 2, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "reversed list",
		input:    []int{5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5},
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

func TestInsertionSort(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			InsertionSort(testCase.input)
			if !reflect.DeepEqual(testCase.input, testCase.expected) {
				t.Errorf("InsertionSort(%v) == %v, want %v",
					testCase.input, testCase.input, testCase.expected)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			SelectionSort(testCase.input)
			if !reflect.DeepEqual(testCase.input, testCase.expected) {
				t.Errorf("SelectionSort(%v) == %v, want %v",
					testCase.input, testCase.input, testCase.expected)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			BubbleSort(testCase.input)
			if !reflect.DeepEqual(testCase.input, testCase.expected) {
				t.Errorf("BubbleSort(%v) == %v, want %v",
					testCase.input, testCase.input, testCase.expected)
			}
		})
	}
}
