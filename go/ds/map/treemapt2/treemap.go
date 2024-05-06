package treemapt2

import cmap "dsa/ds/map"

type treeMap[T any] struct {
	root *node[T]
}

func New[T any]() cmap.Map[string, T] {
	return &treeMap[T]{
		root: newNode[T](),
	}
}

func (t *treeMap[T]) Put(key string, value T) {
	walk := t.root
	for _, r := range key {
		n, exists := walk.children[r]
		if !exists {
			n = newNode[T]()
			walk.children[r] = n
		}
		walk = n
	}
	walk.hasData = true
	walk.data = value
}

func (t *treeMap[T]) Get(key string) (T, bool) {
	walk := t.root
	for _, r := range key {
		n, exists := walk.children[r]
		if !exists {
			return *new(T), false
		}
		walk = n
	}
	return walk.data, walk.hasData
}

func (t *treeMap[T]) Keys() []string {
	var keys []string
	dfs(t.root, nil, func(key string, value T) {
		keys = append(keys, key)
	})
	return keys
}

func (t *treeMap[T]) Values() []T {
	var values []T
	dfs(t.root, nil, func(key string, value T) {
		values = append(values, value)
	})
	return values
}

func (t *treeMap[T]) Entries() []*cmap.Entry[string, T] {
	var entries []*cmap.Entry[string, T]
	dfs(t.root, nil, func(key string, value T) {
		entries = append(entries, cmap.NewEntry(key, value))
	})
	return entries
}

func (t *treeMap[T]) Clear() {
	t.root = newNode[T]()
}

func (t *treeMap[T]) Remove(key string) (T, bool) {
	walk := t.root
	for _, r := range key {
		n, exists := walk.children[r]
		if !exists {
			return *new(T), false
		}
		walk = n
	}
	data := walk.data
	hasData := walk.hasData
	walk.hasData = false
	return data, hasData
}
