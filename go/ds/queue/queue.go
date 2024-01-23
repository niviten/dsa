package queue

import (
	"errors"
	"fmt"
)

type node[T any] struct {
	data T
	next *node[T]
}

type Queue[T any] struct {
	front *node[T]
	rear  *node[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{nil, nil}
}

func (q *Queue[T]) Print() {
	fmt.Print("[")
	walk := q.front
	for walk != nil {
		fmt.Printf("%v, ", walk.data)
		walk = walk.next
	}
	fmt.Println("]")
}

func (q *Queue[T]) Enqueue(data T) {
	newNode := &node[T]{data, nil}
	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
		return
	}
	q.rear.next = newNode
	q.rear = newNode
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.front == nil {
		return *new(T), errors.New("Queue is empty")
	}
	data := q.front.data
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	return data, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.front == nil {
		return *new(T), errors.New("Queue is empty")
	}
	return q.front.data, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return q.front == nil
}
