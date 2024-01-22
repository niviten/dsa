package sort

func MergeSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    mid := len(arr) / 2
    left := arr[:mid]
    right := arr[mid:]
    MergeSort(left)
    MergeSort(right)
    result := merge(left, right)
    copy(arr, result)
}

func merge(a []int, b []int) []int {
    aLen := len(a)
    bLen := len(b)
    result := make([]int, aLen + bLen)
    i := 0
    j := 0
    idx := 0

    for i < aLen && j < bLen {
        if a[i] < b[j] {
            result[idx] = a[i]
            i = i + 1
        } else {
            result[idx] = b[j]
            j = j + 1
        }
        idx = idx + 1
    }

    for i < aLen {
        result[idx] = a[i]
        idx = idx + 1
        i = i + 1
    }

    for j < bLen {
        result[idx] = b[j]
        idx = idx + 1
        j = j + 1
    }

    return result
}
