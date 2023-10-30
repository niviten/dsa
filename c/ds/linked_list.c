#include <stdio.h>
#include <stdlib.h>

typedef struct _Node {
    int data;
    struct _Node *next;
} Node;

typedef struct {
    Node *head;
} LinkedList;

void init(LinkedList*);
void destroy(LinkedList*);
void print(LinkedList*);
void insert(LinkedList*, int);
void remove_element(LinkedList*, int);

void init(LinkedList *list) {
    list->head = NULL;
}

void destroy(LinkedList *list) {
    Node *node;
    Node *temp;

    node = list->head;

    while (node != NULL) {
        temp = node;
        node = node->next;
        free(temp);
    }

    list->head = NULL;
}

void print(LinkedList *list) {
    Node *node;

    printf("[");
    node = list->head;
    while (node != NULL) {
        printf("%d, ", node->data);
        node = node->next;
    }
    printf("]\n");
}

void insert(LinkedList *list, int data) {
    Node *node;
    Node *walk;

    node = (Node*) malloc(sizeof(Node));
    node->data = data;
    if (list->head == NULL) {
        list->head = node;
        return;
    }
    walk = list->head;
    while (walk->next != NULL) {
        walk = walk->next;
    }
    walk->next = node;
}

void remove_element(LinkedList *list, int data) {
    Node *curr;
    Node *temp;

    if (list->head == NULL) {
        return;
    }

    if (list->head->data == data) {
        temp = list->head;
        list->head = list->head->next;
        free(temp);
        return;
    }

    curr = list->head;
    while (curr->next != NULL) {
        if (curr->next->data == data) {
            temp = curr->next;
            curr->next = curr->next->next;
            free(temp);
            break;
        }
        curr = curr->next;
    }
}
