#include <stdio.h>
#include <stdlib.h>

typedef struct {
    int *arr;
    int top;
    int capacity;
} Stack;

void stack_init(Stack *stack) {
    stack->capacity = 10;
    stack->top = 0;
    stack->arr = malloc(sizeof(int) * stack->capacity);
}

void stack_print(Stack *stack) {
    int *iter;
    printf("[");
    for (iter = stack->arr; iter != (stack->arr + stack->top); iter++) {
        printf("%d, ", *iter);
    }
    printf("]\n");
}

void stack_push(Stack *stack, int data) {
    if (stack->top == stack->capacity) {
        stack_increase_capacity(stack);
    }
    *(stack->arr + stack->top) = data;
    stack->top = stack->top + 1;
}

void stack_pop(Stack *stack, int *data) {
    if (stack->top == 0) {
        return;
    }
    *data = *(stack->arr + stack->top - 1);
    stack->top = stack->top - 1;
}

void stack_increase_capacity(Stack *stack) {
    int *arr;
    int *new_arr;
    int *iter;
    int *new_arr_iter;
    int arr_len;

    arr = stack->arr;
    arr_len = stack->top;

    stack->capacity = stack->capacity * 2;
    stack->arr = (int*) malloc(sizeof(int) * stack->capacity);

    new_arr = stack->arr;
    
    iter = arr;
    new_arr_iter = new_arr;
    while (iter != (arr + arr_len)) {
        *new_arr_iter = *iter;
        iter = iter + 1;
        new_arr_iter = new_arr_iter + 1;
    }

    free(arr);
}

void stack_delete(Stack *stack) {
    stack->capacity = 0;
    stack->top = 0;
    if (stack->arr != NULL) {
        free(stack->arr);
        stack->arr = NULL;
    }
}
