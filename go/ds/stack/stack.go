package stack

import (
	"errors"
	"fmt"
)

type node[T any] struct {
	data T
	next *node[T]
}

type Stack[T any] struct {
	top *node[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{top: nil}
}

func (s *Stack[T]) Print() {
	walk := s.top
	fmt.Print("[")
	for walk != nil {
		fmt.Printf("%v, ", walk.data)
		walk = walk.next
	}
	fmt.Println("]")
}

func (s *Stack[T]) Push(data T) {
	newNode := &node[T]{data, s.top}
	s.top = newNode
}

func (s *Stack[T]) Pop() (T, error) {
	if s.top == nil {
		return *new(T), errors.New("Stack underflow")
	}
	dataToReturn := s.top.data
	s.top = s.top.next
	return dataToReturn, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.top == nil {
		return *new(T), errors.New("Stack underflow")
	}
	return s.top.data, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}
