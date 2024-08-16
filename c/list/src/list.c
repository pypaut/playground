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
    l = l;
    index = index;
    return 0;
}

void remove_at(struct list *l, size_t index) {
    l = l;
    index = index;
}

int pop_back(struct list *l) {
    l = l;
    return 0;
}

int pop_front(struct list *l) {
    l = l;
    return 0;
}

void push_back(struct list *l, int element) {
    l = l;
    element = element;
}

void push_front(struct list *l, int element) {
    l = l;
    element = element;
}
