#pragma once

#include <SDL2/SDL.h>
#include <arpa/inet.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

#include "include.h"

#define PORT 8080

const int W = 1920;
const int H = 1080;
const float SCALE = 0.5;

void init_client(int *client_socket_fd);
Uint64 clock_tick(Uint64 current_tick);
void get_dir(const Uint8 *keys, float *dir_x, float *dir_y);
int check_quit_events(const Uint8 *keys);
int draw(SDL_Renderer **renderer);
