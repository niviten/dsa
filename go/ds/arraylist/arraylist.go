package arraylist

import "errors"

type ArrayList[T any] struct {
    arr []T
    size int
    capacity int
}

func New[T any]() *ArrayList[T] {
    return &ArrayList[T]{
        arr: make([]T, 10),
        size: 0,
        capacity: 10,
    }
}

func (list *ArrayList[T]) Insert(item T) {
    list.arr[list.size] = item
    list.size = list.size + 1
    if list.size == list.capacity {
        list.capacity = list.capacity + 10
        newArr := make([]T, list.capacity)
        for i, a := range list.arr {
            newArr[i] = a
        }
        list.arr = newArr
    }
}

func (list *ArrayList[T]) Get(index int) (T, error) {
    if index < 0 || index >= list.size {
        return *new(T), errors.New("Index out of bound")
    }
    return list.arr[index], nil
}

func (list *ArrayList[T]) Remove(equals func(item T) bool) bool {
    i := 0
    for i < list.size {
        if equals(list.arr[i]) {
            break
        }
        i = i + 1
    }
    if i == list.size {
        return false
    }
    for i < list.size - 1 {
        list.arr[i] = list.arr[i+1]
        i = i + 1
    }
    list.size = list.size - 1
    return true
}

func (list *ArrayList[T]) Array() []T {
    arr := make([]T, list.size)
    for i := 0; i < list.size; i++ {
        arr[i] = list.arr[i]
    }
    return arr
}

func (list *ArrayList[T]) Size() int {
    return list.size
}
