#include <stdio.h>
#include <stdlib.h>

struct grid {
    int *numbers;
};

struct grid *new_grid();
void free_grid(struct grid *g);
void print_grid(struct grid *g);
