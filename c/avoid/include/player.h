#pragma once

#include <SDL2/SDL.h>

typedef struct player {
    SDL_Rect *rect;
} player;

player *new_player(int H, int W);
void destroy_player(player *p);
int draw_player(player *p, SDL_Renderer *renderer);
void update_player(player *p, int mouse_x, int mouse_y);
