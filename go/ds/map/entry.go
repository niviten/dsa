package cmap

type Entry[K any, V any] struct {
	Key   K
	Value V
}

func NewEntry[K any, V any](key K, value V) *Entry[K, V] {
	return &Entry[K, V]{
		Key:   key,
		Value: value,
	}
}
