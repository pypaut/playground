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

void update_game(game *g) {
    tick(g->c);

    const Uint8 *keys = SDL_GetKeyboardState(NULL);
    update_player(g->p1, keys, g->H);
    update_player(g->p2, keys, g->H);

    if (!g->has_started && keys[SDL_SCANCODE_SPACE]) {
        g->has_started = 1;
        g->b->dir_x = 1;
    }

    if (g->has_started && update_ball(g->b, g->p1, g->p2, g->H, g->W)) {
        g->is_running = 0;
    }
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

    if (draw_player(g->p1, *g->renderer) || draw_player(g->p2, *g->renderer)) {
        g->error = 1;
    }

    if (draw_ball(g->b, *g->renderer)) {
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
    destroy_player(g->p1);
    destroy_player(g->p2);
    destroy_ball(g->b);
    free(g);

    SDL_Quit();
}
