#include <list.h>

struct list *new_list() {
    struct list *new_list = NULL;
    return new_list;
}

void free_list(struct list *l) {
    if (!l) {
        return;
    }

    struct list *old_head;
    while (l) {
        old_head = l;
        l = l->next;
        free(old_head);
    }
}

int at(struct list *l, size_t index) {
    struct list *current = l;
    for (size_t i = 0; i < index; i++) {
        current = current->next;
    }

    return current->value;
}

int pop_at(struct list *l, size_t index) {
    struct list *current = l;
    int value = 0;

    // Get value
    for (size_t i = 0; i < index; i++) {
        current = current->next;
    }
    value = current->value;

    // Rm list element
    current = l;
    for (size_t i = 0; i < index-1; i++) {
        current = current->next;
    }
    current->next = current->next->next;
    return value;
}

int pop_back(struct list *l) {
    l = l;
    return 0;
}

int pop_front(struct list **l) {
    int value = (*l)->value;

    struct list *old_head;
    struct list *new_head;

    old_head = *l;
    new_head = (*l)->next;

    free(old_head);
    *l = new_head;

    return value;
}

struct list *push_back(struct list *l, int value) {
    l = l;
    value = value;
    return NULL;
}

void push_front(struct list **l, int value) {
    struct list *new_head = calloc(1, sizeof(struct list));
    new_head->next = *l;
    new_head->value = value;
    *l = new_head;
}

size_t len(struct list *l) {
    if (!l) {
        return 0;
    }

    size_t len = 1;
    while (l->next) {
        len++;
        l = l->next;
    }

    return len;
}

void pprint(struct list *l) {
    if (!l) {
        printf("[]\n");
        return;
    }

    char *to_print = calloc(1024, sizeof(char));
    char *to_concat = calloc(1024, sizeof(char));
    strcat(to_print, "[");

    // Head
    sprintf(to_concat, "%d", l->value);
    strcat(to_print, to_concat);
    l = l->next;

    // Tail (with arrows)
    while (l) {
        sprintf(to_concat, " -> %d", l->value);
        strcat(to_print, to_concat);
        l = l->next;
    }

    strcat(to_print, "]");
    printf("%s\n", to_print);

    free(to_concat);
    free(to_print);
}
