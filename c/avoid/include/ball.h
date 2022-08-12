#pragma once

#include <SDL2/SDL.h>

typedef struct ball {
    SDL_Rect *rect;
    float x;
    float y;
    float dir_x;
    float dir_y;
    float speed;
} ball;

ball *new_ball(float x, float y, float dir_x, float dir_y);
void destroy_ball(ball *b);
int draw_ball(ball *b, SDL_Renderer *renderer);
int update_ball(ball *b, int W, int H);
