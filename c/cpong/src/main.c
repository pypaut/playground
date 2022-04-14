#include <player.h>
#include <ball.h>


int main() {
    int W = 1000;
    int H = 800;

    SDL_Init(SDL_INIT_VIDEO|SDL_INIT_AUDIO);

    SDL_Window *window;
    SDL_Renderer *renderer;

    if (SDL_CreateWindowAndRenderer(W, H, 0, &window, &renderer)) {
        fprintf(stderr, "%s\n", "Error on Window/Renderer creation\0");
        return 1;
    }

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

        // Draw
        if (SDL_SetRenderDrawColor(renderer, 0, 0, 0, 255)) {
            fprintf(stderr, "%s\n", "Error Renderer.SetRendererDrawColor\0");
            return 1;
        }

        if (SDL_RenderClear(renderer)) {
            fprintf(stderr, "%s\n", "Error Renderer.RenderClear\0");
            return 1;
        }

        if (draw_player(p1, renderer) || draw_player(p2, renderer)) {
            error = 1;
        }

        if (draw_ball(b, renderer)) {
            error = 1;
        }
        
        SDL_RenderPresent(renderer);
    }

    destroy_player(p1);
    destroy_player(p2);

    SDL_Quit();
    return 0;
}
