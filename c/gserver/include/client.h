#pragma once

#include <SDL2/SDL.h>
#include <arpa/inet.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

#include "constants.h"
#include "extract.h"


void init_client(int *client_socket_fd);
Uint64 clock_tick(Uint64 current_tick);
void get_dir(const Uint8 *keys, float *dir_x, float *dir_y);
int check_quit_events(const Uint8 *keys);
int draw(SDL_Renderer **renderer);
