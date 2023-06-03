#include <grid.h>


struct grid *new_grid() {
    struct grid *g = calloc(1, sizeof(struct grid));
    size_t *numbers = calloc(9*9, sizeof(size_t));
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
                printf("| %ld ", g->numbers[i * 9 + j]);
            }
            else {
                printf("| %ld |", g->numbers[i * 9 + j]);
            }
        }
        printf("\n");
    }
    printf("\n");
}

int compute_grid(struct grid *g) {
    // Detect first 0 position
    size_t pos_to_fill = 0;
    for (size_t i = 0; i < 81; i++) {
        if (g->numbers[i] == 0) {
            pos_to_fill = i;
            break;
        }
        if (i == 80) {
            return 0;
        }
    }

    // Compute line and column
    size_t line = pos_to_fill / 9;
    size_t column = pos_to_fill % 9;

    // Compute first compatible number
    for (size_t nb = 1; nb < 10; nb++) {
        size_t found_another = 0;

        // Check line
        for (size_t j = 0; j < 9; j++) {
            if (g->numbers[line * 9 + j] == nb) {
                found_another = 1;
                break;
            }
        }

        // Check column
        for (size_t i = 0; i < 9; i++) {
            if (g->numbers[i * 9 + column] == nb) {
                found_another = 1;
                break;
            }
        }

        // Check inner grid
        // FIXME

        if (!found_another) {
            g->numbers[pos_to_fill] = nb;
            return compute_grid(g);
        }
    }

    return 1;
}
