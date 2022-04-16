#include <player.h>
#include <ball.h>
#include <game.h>


int main() {
    int W = 1000;
    int H = 800;

    game *g = init_game(W, H);

    int is_running = 1;
    int error = 0;
    int has_started = 0;

    Uint64 last = 0;
    Uint64 now = 0;
    Uint64 dt = 0;

    while (is_running && !error) {
        // Events
        SDL_Event event;
        while (SDL_PollEvent(&event)) {
            if (event.type == SDL_QUIT) {
                is_running = 0;
            }
        }

        const Uint8 *keys = SDL_GetKeyboardState(NULL);

        // Update
        last = now;
        now = SDL_GetTicks();
        dt = now - last;
        if (dt < 1000 / 60) {
            SDL_Delay(1000 / 60 - dt);
        }

        update_game(g, keys);
        if (!has_started && keys[SDL_SCANCODE_SPACE]) {
            has_started = 1;
            g->b->dir_x = 1;
        }

        if (has_started && update_ball(g->b, g->p1, g->p2, g->H, g->W)) {
            is_running = 0;
        }

        // Draw
        if (SDL_SetRenderDrawColor(*g->renderer, 0, 0, 0, 255)) {
            fprintf(stderr, "%s\n", "Error Renderer.SetRendererDrawColor\0");
            return 1;
        }

        if (SDL_RenderClear(*g->renderer)) {
            fprintf(stderr, "%s\n", "Error Renderer.RenderClear\0");
            return 1;
        }

        if (draw_player(g->p1, *g->renderer) || draw_player(g->p2, *g->renderer)) {
            error = 1;
        }

        if (draw_ball(g->b, *g->renderer)) {
            error = 1;
        }
        
        SDL_RenderPresent(*g->renderer);
    }

    destroy_game(g);

    // SDL_Quit();
    return 0;
}
