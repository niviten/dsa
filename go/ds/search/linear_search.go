package search

func LinearSearch[T comparable](haystack []T, needle T) int {
    for idx, a := range haystack {
        if a == needle {
            return idx
        }
    }
    return -1
}

type SearchEquals[T any] func(T, T) bool
func LinearSearchWithComparator[T any](haystack []T, needle T,
        equals SearchEquals[T]) int {
    for idx, a := range haystack {
        if equals(a, needle) {
            return idx
        }
    }
    return -1
}
