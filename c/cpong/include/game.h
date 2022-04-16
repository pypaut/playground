#pragma once

#include <SDL2/SDL.h>
#include <ball.h>
#include <player.h>

typedef struct game {
    SDL_Window **window;
    SDL_Renderer **renderer;
} game;

game *init_game(int W, int H);
void destroy_game(game *g);
