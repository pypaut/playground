#pragma once

#include <SDL2/SDL.h>

typedef struct ball {
    SDL_Rect *rect;
    int r;
    int g;
    int b;
    int a;
    float dir_x;
    float dir_y;
} ball;

ball *new_ball(int x, int y, int w, int h);
void destroy_ball(ball *b);
void set_ball_color(ball *ba, int r, int g, int b, int a);
int draw_ball(ball *b, SDL_Renderer *renderer);
