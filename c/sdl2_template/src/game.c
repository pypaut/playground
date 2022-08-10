#include <game.h>

game *init_game(int W, int H) {
    SDL_Init(SDL_INIT_VIDEO|SDL_INIT_AUDIO);

    game *g = calloc(1, sizeof(game));

    g->H = H;
    g->W = W;

    g->window = calloc(1, 8);
    g->renderer = calloc(1, 8);

    g->c = new_clock();

    g->is_running = 1;
    g->error = 0;
    g->has_started = 0;

    if (SDL_CreateWindowAndRenderer(W, H, 0, g->window, g->renderer)) {
        fprintf(stderr, "%s\n", "Error on Window/Renderer creation\0");
        return NULL;
    }

    return g;
}

void update_game(game *g) {
    tick(g->c);

    const Uint8 *keys = SDL_GetKeyboardState(NULL);
    keys = keys;
}

void draw_game(game *g) {
    if (SDL_SetRenderDrawColor(*g->renderer, 0, 0, 0, 255)) {
        fprintf(stderr, "%s\n", "Error Renderer.SetRendererDrawColor\0");
        g->error = 1;
    }

    if (SDL_RenderClear(*g->renderer)) {
        fprintf(stderr, "%s\n", "Error Renderer.RenderClear\0");
        g->error = 1;
    }

    SDL_RenderPresent(*g->renderer);
}

void handle_quit_event(game *g) {
    SDL_Event event;
    while (SDL_PollEvent(&event)) {
        if (event.type == SDL_QUIT) {
            g->is_running = 0;
        }
    }
}

void destroy_game(game *g) {
    free(g->window);
    free(g->renderer);
    destroy_clock(g->c);
    free(g);

    SDL_Quit();
}
