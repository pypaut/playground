#include <grid.h>

int main(void) {
    struct grid *g = new_grid();
    print_grid(g);
    free_grid(g);
    return 0;
}
