package queue

import (
    "fmt"
)

type node struct {
    data int
    next *node
}

type Queue struct {
    front *node
    rear *node
}

func NewQueue() *Queue {
    return &Queue{nil, nil}
}

func (q *Queue) Print() {
    fmt.Print("[")
    walk := q.front
    for walk != nil {
        fmt.Printf("%d, ", walk.data)
        walk = walk.next
    }
    fmt.Println("]")
}

func (q *Queue) Enqueue(data int) {
    newNode := &node{data, nil}
    if q.rear == nil {
        q.front = newNode
        q.rear = newNode
        return
    }
    q.rear.next = newNode
    q.rear = newNode
}

func (q *Queue) Dequeue() (int, bool) {
    if q.front == nil {
        return 0, false
    }
    data := q.front.data
    q.front = q.front.next
    if q.front == nil {
        q.rear = nil
    }
    return data, true
}

func (q *Queue) Peek() (int, bool) {
    if q.front == nil {
        return 0, false
    }
    return q.front.data, true
}

func (q *Queue) IsEmpty() bool {
    return q.front == nil
}
