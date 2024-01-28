package sort

func QuickSort(arr []int) {
    quickSort(arr, 0, len(arr) - 1)
}

func quickSort(arr []int, low int, high int) {
    if low >= high || low < 0 {
        return
    }

    p := partition(arr, low, high)

    quickSort(arr, low, p - 1)
    quickSort(arr, p + 1, high)
}

func partition(arr[] int, low int, high int) int {
    pivot := arr[high]
    
    i := low - 1

    for j := low; j <= high-1; j++ {
        if arr[j] <= pivot {
            i = i + 1
            tmp := arr[i]
            arr[i] = arr[j]
            arr[j] = tmp
        }
    }

    i = i + 1
    tmp := arr[i]
    arr[i] = arr[high]
    arr[high] = tmp

    return i
}
