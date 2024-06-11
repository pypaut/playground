#pragma once

#include <arpa/inet.h>
#include <errno.h>
#include <netinet/in.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/select.h>
#include <sys/socket.h>
#include <sys/time.h>
#include <unistd.h>

#include "constants.h"
#include "extract.h"

double clamp(double d, double min, double max);
void add_child_sockets_to_set(fd_set *readfds, int *client_sockets, int *max_sd);
void handle_clients_messages(int *client_sockets, fd_set *readfds, struct sockaddr_in *address, float *pos_x, float *pos_y);
void handle_new_client(int *new_socket, int *master_socket, int *client_sockets, struct sockaddr_in *address);
void init_server(int *server_socket_fd, struct sockaddr_in *address);
void update_pos(float *pos_x, float *pos_y, float dir_x, float dir_y);
