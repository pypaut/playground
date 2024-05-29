#pragma once

#include <SDL2/SDL.h>

typedef struct player {
    double x;
    double y;
    
    int w;
    int h;

    SDL_Rect *rect;
    SDL_Color *color;
} player;

player *create_player();
void destroy_player(player *p);
void update_player(player *p);
void draw_player(player *p);
void update_player_rect(player *p);
