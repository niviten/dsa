package search

func LinearSearch[T comparable](haystack []T, needle T) int {
	for idx, a := range haystack {
		if a == needle {
			return idx
		}
	}
	return -1
}

func LinearSearchWithComparator[T any](haystack []T, needle T,
	compare Compare[T]) int {
	for idx, a := range haystack {
		if compare(a, needle) == 0 {
			return idx
		}
	}
	return -1
}
