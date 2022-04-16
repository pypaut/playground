#include <game.h>

game *init_game(int W, int H) {
    SDL_Init(SDL_INIT_VIDEO|SDL_INIT_AUDIO);

    game *g = calloc(1, sizeof(game));

    g->window = calloc(1, 8);
    g->renderer = calloc(1, 8);

    if (SDL_CreateWindowAndRenderer(W, H, 0, g->window, g->renderer)) {
        fprintf(stderr, "%s\n", "Error on Window/Renderer creation\0");
        return NULL;
    }

    return g;
}

void destroy_game(game *g) {
    free(g->window);
    free(g->renderer);
    free(g);

    SDL_Quit();
}
