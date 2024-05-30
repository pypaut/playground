#include <game.h>
#include <clock.h>


int main() {
    int W = 1920 / 2;
    int H = 1080 / 2;

    game *g = init_game(W, H);

    while (g->is_running && !g->error) {
        handle_quit_event(g);
        update_game(g);
        draw_game(g);
    }

    destroy_game(g);
    return 0;
}
