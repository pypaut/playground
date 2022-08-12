#pragma once

#include <ball.h>
#include <clock.h>
#include <player.h>

typedef struct game {
    SDL_Window **window;
    SDL_Renderer **renderer;

    clock *c;
    player *p;

    ball **balls;
    int nb_balls;

    int H;
    int W;

    int is_running;
    int error;
    int has_started;
} game;

game *init_game();
void destroy_game(game *g);
void update_game(game *g);
void draw_game(game *g);
void handle_quit_event(game *g);
void spawn_ball(game *g, int pos);
