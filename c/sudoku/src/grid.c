#include <grid.h>


struct grid *new_grid() {
    struct grid *g = calloc(1, sizeof(struct grid));
    int *numbers = calloc(9*9, sizeof(int));
    if (!g || !numbers) {
        return NULL;
    }

    g->numbers = numbers;
    return g;
}

void free_grid(struct grid *g) {
    free(g->numbers);
    free(g);
}

void print_grid(struct grid *g) {
    for (size_t i = 0; i < 9; i++) {
        for (size_t j = 0; j < 9; j++) {
            if (j != 8) {
                printf("| %d ", g->numbers[i * 9 + j]);
            }
            else {
                printf("| %d |", g->numbers[i * 9 + j]);
            }
        }
        printf("\n");
    }
}
