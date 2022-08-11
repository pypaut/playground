#include <game.h>
#include <clock.h>


int main() {
    int H = 800;
    int W = 1000;

    game *g = init_game(H, W);

    while (g->is_running && !g->error) {
        handle_quit_event(g);
        update_game(g);
        draw_game(g);
    }

    destroy_game(g);
    return 0;
}
