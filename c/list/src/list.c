#include <list.h>

struct list *new_list() {
    struct list *new_list = calloc(1, sizeof(struct list));
    return new_list;
}

void free_list(struct list *l) {
    struct list *prev;

    while (l->next) {
        prev = l;
        l = l->next;
        free(prev);
    }
}

int at(struct list *l, size_t index) {
    struct list *current = l;
    for (size_t i = 0; i < index; i++) {
        current = current->next;
    }

    return current->value;
}

void remove_at(struct list *l, size_t index) {
    l = l;
    index = index;
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

struct list *push_front(struct list *l, int value) {
    struct list *new_head = new_list();
    new_head->next = l;
    new_head->value = value;
    return new_head;
}
