#include <game.h>

game *init_game(int W, int H) {
    SDL_Init(SDL_INIT_VIDEO|SDL_INIT_AUDIO);

    game *g = calloc(1, sizeof(game));

    g->H = H;
    g->W = W;

    g->window = calloc(1, 8);
    g->renderer = calloc(1, 8);

    g->p1 = new_player(100, H/2 - 50, 10, 100, 1);
    set_player_color(g->p1, 150, 0, 150, 255);

    g->p2 = new_player(W - 110, H/2 - 50, 10, 100, 2);
    set_player_color(g->p2, 150, 0, 150, 255);
    
    g->b = new_ball(W/2 - 5, H/2 - 5, 10, 10);
    set_ball_color(g->b, 255, 255, 255, 255);

    if (SDL_CreateWindowAndRenderer(W, H, 0, g->window, g->renderer)) {
        fprintf(stderr, "%s\n", "Error on Window/Renderer creation\0");
        return NULL;
    }

    return g;
}

void update_game(game *g, const Uint8 *keys) {
    update_player(g->p1, keys, g->H);
    update_player(g->p2, keys, g->H);
}

void destroy_game(game *g) {
    free(g->window);
    free(g->renderer);
    destroy_player(g->p1);
    destroy_player(g->p2);
    destroy_ball(g->b);
    free(g);

    SDL_Quit();
}
