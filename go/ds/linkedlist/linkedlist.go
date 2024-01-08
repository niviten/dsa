package linkedlist

import (
	"fmt"
)

type Node[T any] struct {
	data T
	next *Node[T]
}

type NodeEquals[T any] func(T, T) bool

func (node *Node[T]) Data() T {
	return node.data
}

type LinkedList[T any] struct {
	head *Node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{head: nil}
}

func (list *LinkedList[T]) Print() {
	fmt.Println("[")
	walk := list.head
	for walk != nil {
		fmt.Println(walk.data)
		walk = walk.next
	}
	fmt.Println("]")
	fmt.Println("___")
}

func (list *LinkedList[T]) Insert(data T) {
	newNode := &Node[T]{data: data, next: nil}
	if list.head == nil {
		list.head = newNode
		return
	}
	walk := list.head
	for walk.next != nil {
		walk = walk.next
	}
	walk.next = newNode
}

func (list *LinkedList[T]) Delete(data T, equals NodeEquals[T]) {
	var prewalk *Node[T] = nil
	walk := list.head

	for walk != nil && !equals(walk.data, data) {
		prewalk = walk
		walk = walk.next
	}

	if walk == nil {
		return
	}

	if prewalk == nil {
		list.head = list.head.next
		return
	}

	prewalk.next = walk.next
}

func (list *LinkedList[T]) Search(searchKey T, equals NodeEquals[T]) int {
	walk := list.head
	index := 0
	for walk != nil {
		// if walk.data == searchKey {
		if equals(walk.data, searchKey) {
			return index
		}
		index = index + 1
		walk = walk.next
	}
	return -1
}

func (list *LinkedList[T]) Length() int {
	length := 0
	walk := list.head
	for walk != nil {
		walk = walk.next
		length = length + 1
	}
	return length
}

func (list *LinkedList[T]) Reverse() *LinkedList[T] {
	walk := list.head
	var n *Node[T]
	n = nil
	for walk != nil {
		newNode := &Node[T]{data: walk.data, next: n}
		n = newNode
		walk = walk.next
	}
	newList := NewLinkedList[T]()
	newList.head = n
	return newList
}

func (list *LinkedList[T]) ReverseInplace() {
	var curr *Node[T]
	var prev *Node[T]
	var temp *Node[T]

	curr = list.head
	prev = nil

	for curr != nil {
		temp = curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}

	list.head = prev
}

func (list *LinkedList[T]) HasCycle() bool {
	if list.head == nil {
		return false
	}

	slow := list.head
	fast := list.head.next

	for fast != nil && fast.next != nil {
		if slow == fast {
			return true
		}
		slow = slow.next
		fast = fast.next.next
	}

	return false
}

func (list *LinkedList[T]) ForEach(callback func(T, int)) {
	walk := list.head
	index := 0
	for walk != nil {
		callback(walk.data, index)
		walk = walk.next
		index = index + 1
	}
}

func (list *LinkedList[T]) Map(callback func(T, int) T) *LinkedList[T] {
	newList := NewLinkedList[T]()
	walk := list.head
	var node *Node[T]
	index := 0

	for walk != nil {
		newNode := &Node[T]{callback(walk.data, index), nil}
		if newList.head == nil {
			newList.head = newNode
			node = newNode
		} else {
			node.next = newNode
			node = node.next
		}
		walk = walk.next
		index = index + 1
	}

	return newList
}

func (list *LinkedList[T]) Filter(callback func(T, int) bool) *LinkedList[T] {
	newList := NewLinkedList[T]()
	walk := list.head
	var node *Node[T]
	index := 0

	for walk != nil {
		canInsert := callback(walk.data, index)
		if canInsert {
			newNode := &Node[T]{walk.data, nil}
			if newList.head == nil {
				newList.head = newNode
				node = newNode
			} else {
				node.next = newNode
				node = node.next
			}
		}
		walk = walk.next
	}
	return newList
}
