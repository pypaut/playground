#pragma once

#include <netinet/in.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>
#define PORT 8080

void init_server(int *server_socket_fd, struct sockaddr_in *address);
void extract_dir(char *buffer, float *dir_x, float *dir_y);
