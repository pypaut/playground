#include <stdio.h>
#include <stdlib.h>

struct grid {
    size_t *numbers;
};

struct grid *new_grid();
void free_grid(struct grid *g);
void print_grid(struct grid *g);
int compute_grid(struct grid *g);
