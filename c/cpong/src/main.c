#include <game.h>
#include <clock.h>


int main() {
    int W = 1000;
    int H = 800;

    game *g = init_game(W, H);

    while (g->is_running && !g->error) {
        handle_quit_event(g);

        const Uint8 *keys = SDL_GetKeyboardState(NULL);
        update_game(g, keys);
        draw_game(g);
    }

    destroy_game(g);
    return 0;
}
