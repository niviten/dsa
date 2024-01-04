package linkedlist

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func (node *Node) Data() int {
	return node.data
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil}
}

func (list *LinkedList) Print() {
	fmt.Print("[")
	walk := list.head
	for walk != nil {
		fmt.Printf("%d, ", walk.data)
		walk = walk.next
	}
	fmt.Println("]")
}

func (list *LinkedList) Insert(data int) {
	newNode := &Node{data: data, next: nil}
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

func (list *LinkedList) Delete(data int) {
	var prewalk *Node = nil
	walk := list.head

	for walk != nil && walk.data != data {
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

func (list *LinkedList) Search(searchKey int) int {
	walk := list.head
	index := 0
	for walk != nil {
		if walk.data == searchKey {
			return index
		}
		index = index + 1
		walk = walk.next
	}
	return -1
}

func (list *LinkedList) Length() int {
	length := 0
	walk := list.head
	for walk != nil {
		walk = walk.next
		length = length + 1
	}
	return length
}

func (list *LinkedList) Reverse() *LinkedList {
	walk := list.head
	var n *Node
	n = nil
	for walk != nil {
		newNode := &Node{data: walk.data, next: n}
		n = newNode
		walk = walk.next
	}
	newList := NewLinkedList()
	newList.head = n
	return newList
}

func (list *LinkedList) ReverseInplace() {
	var curr *Node
	var prev *Node
	var temp *Node

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

func (list *LinkedList) HasCycle() bool {
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
