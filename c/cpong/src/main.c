#include <game.h>


int main() {
    int W = 1000;
    int H = 800;

    game *g = init_game(W, H);

    Uint64 last = 0;
    Uint64 now = 0;
    Uint64 dt = 0;

    while (g->is_running && !g->error) {
        handle_quit_event(g);

        // Clock
        last = now;
        now = SDL_GetTicks();
        dt = now - last;
        if (dt < 1000 / 60) {
            SDL_Delay(1000 / 60 - dt);
        }

        const Uint8 *keys = SDL_GetKeyboardState(NULL);
        update_game(g, keys);
        draw_game(g);
    }

    destroy_game(g);
    return 0;
}
