#pragma once

#include <player.h>

typedef struct ball {
    SDL_Rect *rect;
    float x;
    float y;
    int r;
    int g;
    int b;
    int a;
    float dir_x;
    float dir_y;
    float speed;
} ball;

ball *new_ball(float x, float y, int w, int h);
void destroy_ball(ball *b);
void set_ball_color(ball *ba, int r, int g, int b, int a);
int draw_ball(ball *b, SDL_Renderer *renderer);
int update_ball (ball *b, player *p1, player *p2, int H, int W);
