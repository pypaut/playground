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

    if (SDL_SetWindowFullscreen(*g->window, SDL_WINDOW_FULLSCREEN_DESKTOP)) {
        fprintf(stderr, "%s\n", "Error on window fullscreen\0");
        return NULL;
    }
    SDL_GetWindowSize(*g->window, &g->W, &g->H);

    SDL_ShowCursor(SDL_DISABLE);

    g->c = new_clock();
    g->p = new_player(g->H, g->W);
    g->balls = calloc(1, 8);
    g->nb_balls = 0;

    // Spawn one ball (for now)
    spawn_ball(g, 1);

    g->is_running = 1;
    g->error = 0;
    g->has_started = 0;

    return g;
}

void update_game(game *g) {
    tick(g->c);

    int x = 0;
    int y = 0;

    SDL_GetMouseState(&x, &y);
    update_player(g->p, x, y);
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

    draw_player(g->p, *g->renderer);

    for (int i = 0; i < g->nb_balls; i++) {
        draw_ball(g->balls[i], *g->renderer);
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
    destroy_player(g->p);
    for (int i = 0; i < g->nb_balls; i++) {
        destroy_ball(g->balls[i]);
    }
    free(g->balls);
    free(g);

    SDL_Quit();
}

void spawn_ball(game *g, int pos) {
    float x, y, dir_x, dir_y;
    if (pos == 1) {  // Top left corner
        x = 0;
        y = 0;
        dir_x = 0.5;
        dir_y = 0.5;
    }
    else if (pos == 2) {  // Top right corner

    }
    else if (pos == 3) {  // Bottom left corner

    }
    else {  // Bottom right corner

    }

    ball *b = new_ball(x, y, dir_x, dir_y);

    // Update game struct
    g->nb_balls += 1;
    g->balls = realloc(g->balls, 8 * g->nb_balls);
    g->balls[g->nb_balls-1] = b;
}
