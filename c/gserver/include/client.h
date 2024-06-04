#pragma once

#include <SDL2/SDL.h>
#include <arpa/inet.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

#define PORT 8080

void init_client(int *client_socket_fd);
Uint64 clock_tick(Uint64 current_tick);
void get_dir(const Uint8 *keys, float *dir_x, float *dir_y);
