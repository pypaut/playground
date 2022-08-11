#include <game.h>
#include <clock.h>


int main() {
    game *g = init_game();

    while (g->is_running && !g->error) {
        handle_quit_event(g);
        update_game(g);
        draw_game(g);
    }

    destroy_game(g);
    return 0;
}
