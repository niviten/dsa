package sort

func InsertionSort(arr []int) {
	arrLen := len(arr)
	for i := 1; i < arrLen; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}
