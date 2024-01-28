#include <stdio.h>
#include <stdlib.h>

typedef struct _Node {
    int data;
    struct _Node *prev;
    struct _Node *next;
} Node;

typedef struct {
    Node *head;
} LinkedList;

typedef LinkedList List;

void init(List*);
void destroy(List*);
void print(List*);
void print_both_ways(List*);
void insert(List*, int);
void insert_at_beginning(List*, int);
void insert_at_end(List*, int);
void insert_at_position(List*, int, int);
void remove_element(List*, int);
void remove_at_beginning(List*, int*);
void remove_at_end(List*, int*);

void init(List *list) {
    list->head = NULL;
}

void destroy(List *list) {
    Node *walk;
    Node *node;

    walk = list->head;
    while (walk != NULL) {
        node = walk;
        walk = walk->next;
        free(node);
    }
}

void print(List *list) {
    Node *walk;

    walk = list->head;
    printf("[");
    while (walk != NULL) {
        printf("%d, ", walk->data);
        walk = walk->next;
    }
    printf("]\n");
}

void print_both_ways(List *list) {
    Node *walk;
    Node *rev_walk;

    walk = list->head;
    rev_walk = NULL;
    printf("straight: [");
    while (walk != NULL) {
        printf("%d, ", walk->data);
        rev_walk = walk;
        walk = walk->next;
    }
    printf("]\n");

    printf("reverse: [");
    while (rev_walk != NULL) {
        printf("%d, ", rev_walk->data);
        rev_walk = rev_walk->prev;
    }
    printf("]\n");
}

void insert(List *list, int data) {
    Node *node;
    Node *walk;

    node = (Node*) malloc(sizeof(List));
    node->data = data;
    node->next = NULL;
    node->prev = NULL;

    if (list->head == NULL) {
        list->head = node;
        return;
    }

    walk = list->head;
    while (walk->next != NULL) {
        walk = walk->next;
    }
    walk->next = node;
    node->prev = walk;
}


void insert_at_beginning(List *list, int data) {
    Node *node;

    node = (Node*) malloc(sizeof(List));
    node->data = data;
    node->next = list->head;
    node->prev = NULL;

    if (list->head != NULL) {
        list->head->prev = node;
    }

    list->head = node;
}

void insert_at_end(List *list, int data) {
    insert(list, data);
}

void insert_at_position(List *list, int data, int position) {
    int i;
    Node *walk;
    Node *node;

    if (position == 0) {
        insert_at_beginning(list, data);
        return;
    }

    i = 1;
    walk = list->head;
    while (walk != NULL && i < position) {
        walk = walk->next;
        i = i + 1;
    }

    if (walk == NULL) {
        return;
    }

    node = (Node*) malloc(sizeof(Node));
    node->data = data;
    node->next = walk->next;
    node->prev = walk;
    if (walk->next != NULL) {
        walk->next->prev = node;
    }
    walk->next = node;
}

void remove_element(List *list, int data) {
    Node *node;
    Node *prev_node;

    node = list->head;
    prev_node = NULL;
    while (node != NULL && node->data != data) {
        prev_node = node;
        node = node->next;
    }

    if (node == NULL) {
        return;
    }

    if (node == list->head) {
        list->head = list->head->next;
        if (list->head != NULL) {
            list->head->prev = NULL;
        }
        free(node);
        return;
    }

    prev_node->next = node->next;
    if (node->next != NULL) {
        node->next->prev = prev_node;
    }
    free(node);
}

void remove_at_beginning(List *list, int *data) {
    Node *node;

    if (list->head == NULL) {
        return;
    }

    node = list->head;
    list->head = list->head->next;
    if (list->head != NULL) {
        list->head->prev = NULL;
    }
    *data = node->data;
    free(node);
}

void remove_at_end(List *list, int *data) {
    Node *node;

    if (list->head == NULL) {
        return;
    }

    node = list->head;
    while (node->next != NULL) {
        node = node->next;
    }

    if (node == list->head) {
        list->head = NULL;
    } else {
        node->prev->next = NULL;
    }

    *data = node->data;
    free(node);
}
