package search

func BinarySearch[T any](haystack []T, needle T, compare Compare[T]) int {
    haystackLen := len(haystack)
    left := 0
    right := haystackLen - 1

    for left <= right {
        mid := (left + right) / 2
        switch diff := compare(needle, haystack[mid]); diff {
        case 0:
            return mid
        case 1:
            left = mid + 1
        case -1:
            right = mid - 1
        }
    }

    return -1
}
