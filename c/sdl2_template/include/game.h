#pragma once

#include <clock.h>

typedef struct game {
    SDL_Window **window;
    SDL_Renderer **renderer;

    clock *c;

    int H;
    int W;

    int is_running;
    int error;
    int has_started;
} game;

game *init_game(int W, int H);
void destroy_game(game *g);
void update_game(game *g);
void draw_game(game *g);
void handle_quit_event(game *g);
