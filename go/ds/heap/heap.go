package heap

import "fmt"

type Heap struct {
	arr      []int
	size     int
	capacity int
}

func New() *Heap {
	return &Heap{
		arr:      make([]int, 10),
		size:     0,
		capacity: 10,
	}
}

func (h *Heap) Print() {
	fmt.Println(h.arr)
}

func (h *Heap) Insert(item int) {
	idx := h.size
	h.arr[idx] = item

	for parentIdx := parent(idx); parentIdx >= 0; parentIdx = parent(parentIdx) {
		if h.arr[idx] <= h.arr[parentIdx] {
			break
		}
		temp := h.arr[idx]
		h.arr[idx] = h.arr[parentIdx]
		h.arr[parentIdx] = temp
		idx = parentIdx
	}

	h.size = h.size + 1

	if h.size >= h.capacity {
		h.capacity = h.capacity * 2
		newArr := make([]int, h.capacity)
		for i, a := range h.arr {
			newArr[i] = a
		}
		h.arr = newArr
	}
}

func (h *Heap) Extract() (int, bool) {
	if h.size <= 0 {
		return 0, false
	}
	rootElement := h.arr[0]

	h.arr[0] = h.arr[h.size-1]
	h.size = h.size - 1

	idx := 0

	for {
		leftChildIdx := leftChild(idx)
		rightChildIdx := rightChild(idx)
		largest := idx

		if leftChildIdx < h.size && h.arr[leftChildIdx] > h.arr[largest] {
			largest = leftChildIdx
		}

		if rightChildIdx < h.size && h.arr[rightChildIdx] > h.arr[largest] {
			largest = rightChildIdx
		}

		if largest == idx {
			break
		}

		h.arr[idx], h.arr[largest] = h.arr[largest], h.arr[idx]
		idx = largest
	}

	return rootElement, true
}

func parent(index int) int {
	return (index - 1) / 2
}

func leftChild(index int) int {
	return rightChild(index) - 1
}

func rightChild(index int) int {
	return (index + 1) * 2
}
