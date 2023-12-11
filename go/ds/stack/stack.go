package stack

import (
    "fmt"
)

type node struct {
    data int
    next *node
}

type Stack struct {
    top *node
}

func NewStack() *Stack {
    return &Stack{top: nil}
}

func (s *Stack) Print() {
    walk := s.top
    fmt.Print("[")
    for walk != nil {
        fmt.Printf("%d, ", walk.data)
        walk = walk.next
    }
    fmt.Println("]")
}

func (s *Stack) Push(data int) {
    newNode := &node{data, s.top}
    s.top = newNode
}

func (s *Stack) Pop() (int, bool) {
    if s.top == nil {
        return 0, false
    }
    dataToReturn := s.top.data
    s.top = s.top.next
    return dataToReturn, true
}

func (s *Stack) Peek() (int, bool) {
    if s.top == nil {
        return 0, false
    }
    return s.top.data, true
}
