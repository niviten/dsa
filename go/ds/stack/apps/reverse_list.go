package apps

import (
	"dsa/ds/stack"
)

func ReverseList[T any](list []T) error {
	s := stack.NewStack[T]()
	for _, item := range list {
		s.Push(item)
	}
	idx := 0
	for !s.IsEmpty() {
		item, err := s.Pop()
		if err != nil {
			return err
		}
		list[idx] = item
		idx = idx + 1
	}
	return nil
}
