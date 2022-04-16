#include <player.h>
#include <ball.h>
#include <game.h>


int main() {
    int W = 1000;
    int H = 800;

    game *g = init_game(W, H);

    player *p1 = new_player(100, H/2 - 50, 10, 100, 1);
    set_player_color(p1, 150, 0, 150, 255);

    player *p2 = new_player(W - 110, H/2 - 50, 10, 100, 2);
    set_player_color(p2, 150, 0, 150, 255);

    ball *b = new_ball(W/2 - 5, H/2 - 5, 10, 10);
    set_ball_color(b, 255, 255, 255, 255);

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

        update_player(p1, keys, H);
        update_player(p2, keys, H);
        if (!has_started && keys[SDL_SCANCODE_SPACE]) {
            has_started = 1;
            b->dir_x = 1;
        }

        if (has_started && update_ball(b, p1, p2, H, W)) {
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

        if (draw_player(p1, *g->renderer) || draw_player(p2, *g->renderer)) {
            error = 1;
        }

        if (draw_ball(b, *g->renderer)) {
            error = 1;
        }
        
        SDL_RenderPresent(*g->renderer);
    }

    destroy_player(p1);
    destroy_player(p2);
    destroy_ball(b);
    destroy_game(g);

    // SDL_Quit();
    return 0;
}
