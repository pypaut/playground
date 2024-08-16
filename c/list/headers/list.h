#pragma once

#include <stddef.h>
#include <stdlib.h>

struct list {
    int value;
    struct list *next;
};

struct list *new_list();
void free_list(struct list *l);

int at(struct list *l, size_t index);
int pop_at(struct list *l, size_t index);

int pop_back(struct list *l);
int pop_front(struct list *l);

void push_back(struct list *l, int);
void push_front(struct list *l, int);

size_t len(struct list *l);

int max(struct list *l);
int min(struct list *l);

void map(struct list *l, int (*f)(int));
