#pragma once

#include <SDL2/SDL.h>

typedef struct player {
    SDL_Rect *rect;
    int r;
    int g;
    int b;
    int a;
} player;

player *new_player(int x, int y, int w, int h);
void destroy_player(player *p);
void set_player_color(player *p, int r, int g, int b, int a);
int draw_player(player *p, SDL_Renderer *renderer);
