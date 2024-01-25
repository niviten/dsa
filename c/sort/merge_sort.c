#include <stdio.h>
#include <stdlib.h>

void merge_sort(int*, int, int);
void merge(int*, int, int);

int main() {
    int arr[] = {9, 8, 7, 6, 5, 4, 3, 2, 1};
    int arr_len = 9;
    int i;
    merge_sort(arr, 0, arr_len - 1);
    for (i = 0; i < arr_len; i++) {
        printf("%d, ", arr[i]);
    }
    printf("\n");
    printf("___done\n");
    return 0;
}

void merge_sort(int *arr, int from, int to) {
    int mid;
    if ((to - from + 1) <= 1) {
        return;
    }
    mid = from + ((to - from) / 2);
    merge_sort(arr, from, mid);
    merge_sort(arr, mid+1, to);
    merge(arr, from, to);
}

void merge(int *arr, int from, int to) {
    int mid;
    int *left;
    int left_len;
    int *right;
    int right_len;
    int i;
    int j;
    int k;

    mid = from + ((to - from) / 2);
    left_len = mid - from + 1;
    right_len = to - mid;

    left = (int*) malloc(sizeof(int) * left_len);
    right = (int*) malloc(sizeof(int) * right_len);

    for (i = 0; i < left_len; i++) {
        left[i] = arr[from + i];
    }
    for (i = 0; i < right_len; i++) {
        right[i] = arr[mid + 1 + i];
    }

    i = 0;
    j = 0;
    k = from;

    while (i < left_len && j < right_len) {
        if (left[i] < right[j]) {
            arr[k++] = left[i++];
        } else {
            arr[k++] = right[j++];
        }
    }
    while (i < left_len) {
        arr[k++] = left[i++];
    }
    while (j < right_len) {
        arr[k++] = right[j++];
    }

    free(left);
    free(right);
}
