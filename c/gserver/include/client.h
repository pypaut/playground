#pragma once

#include <SDL2/SDL.h>
#include <arpa/inet.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

#include "constants.h"
#include "extract.h"


Uint64 clock_tick(Uint64 current_tick);
int check_quit_events(const Uint8 *keys);
int draw(SDL_Renderer **renderer, SDL_Rect *rect, SDL_Color *color);
void get_dir(const Uint8 *keys, float *dir_x, float *dir_y);
void init_client(int *client_socket_fd);
