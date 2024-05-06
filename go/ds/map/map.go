package cmap

type Map[K any, V any] interface {
	Put(K, V)
	Get(K) (V, bool)
	Keys() []K
	Entries() []*Entry[K, V]
	Clear()
	Remove(K) (V, bool)
}
