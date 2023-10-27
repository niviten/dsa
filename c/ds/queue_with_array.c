#include <stdio.h>
#include <stdlib.h>

typedef struct {
    int *arr;
    int capacity;
    int length;
    int front;
    int rear;
} Queue;

void queue_init(Queue*);
void queue_print(Queue*);
int queue_length(Queue*);
void queue_increase_capacity(Queue*);
void queue_enqueue(Queue*, int);
void queue_dequeue(Queue*, int*);

void queue_init(Queue *stack) {
    stack->capacity = 5;
    stack->arr = (int*) malloc(sizeof(int) * stack->capacity);
    stack->length = 0;
    stack->front = -1;
    stack->rear = -1;
}

void queue_print(Queue *q) {
    int i;
    int index;

    printf("[");
    for (i = 0; i < q->length; i++) {
        index = q->front + i;
        index = index % q->capacity;
        printf("%d, ", q->arr[index]);
    }
    printf("]\n");
}

void queue_increase_capacity(Queue *q) {
    int *new_arr;
    int *new_arr_iter;
    int i;
    int index;
    int old_capacity;
    int *old_arr;

    old_capacity = q->capacity;
    q->capacity = q->capacity * 2;

    new_arr = (int*) malloc(sizeof(int) * q->capacity);

    new_arr_iter = new_arr;
    for (i = 0; i < q->length; i++) {
        index = q->front + i;
        index = index % old_capacity;
        *(new_arr_iter) = q->arr[index];
        new_arr_iter = new_arr_iter + 1;
    }
    old_arr = q->arr;
    q->arr = new_arr;
    free(old_arr);
}

void queue_enqueue(Queue *q, int data) {
    if (q->length == 0) {
        q->front = 0;
    } else if (q->length == q->capacity) {
        queue_increase_capacity(q);
    }
    q->rear = q->rear + 1;
    q->rear = q->rear % q->capacity;
    q->arr[q->rear] = data;
    q->length = q->length + 1;
}

void queue_dequeue(Queue *q, int *out_data) {
    if (q->length == 0) {
        return;
    }
    *out_data = q->arr[q->front];
    q->length = q->length - 1;
    if (q->length == 0) {
        q->front = -1;
        q->rear = -1;
        return;
    }
    q->front = q->front + 1;
    q->front = q->front % q->capacity;
}
