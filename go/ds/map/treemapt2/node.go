package treemapt2

type node[T any] struct {
	data     T
	hasData  bool
	children map[rune]*node[T]
}

func newNode[T any]() *node[T] {
	return &node[T]{
		hasData:  false,
		children: make(map[rune]*node[T]),
	}
}

func dfs[T any](n *node[T], runes []rune, cb func(string, T)) {
	if n.hasData {
		cb(string(runes), n.data)
	}
	for r, child := range n.children {
		runesCopy := make([]rune, len(runes))
		copy(runesCopy, runes)
		runesCopy = append(runesCopy, r)
		dfs(child, runesCopy, cb)
	}
}
