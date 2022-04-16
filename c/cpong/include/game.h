#pragma once

#include <SDL2/SDL.h>
#include <ball.h>
#include <player.h>

typedef struct game {
    SDL_Window **window;
    SDL_Renderer **renderer;

    player *p1;
    player *p2;
    ball *b;

    int H;
    int W;

    int is_running;
    int error;
    int has_started;
} game;

game *init_game(int W, int H);
void destroy_game(game *g);
void update_game(game *g, const Uint8 *keys);
void draw_game(game *g);
