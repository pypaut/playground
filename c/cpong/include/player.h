#pragma once

#include <SDL2/SDL.h>

typedef struct player {
    SDL_Rect *rect;
    int r;
    int g;
    int b;
    int a;
    float speed;
    int nb;
} player;

player *new_player(int x, int y, int w, int h, int nb);
void destroy_player(player *p);
void set_player_color(player *p, int r, int g, int b, int a);
int draw_player(player *p, SDL_Renderer *renderer);
void update_player(player *p, const Uint8 *keys, int H);
