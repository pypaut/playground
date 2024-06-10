#pragma once

#define MAX_CLIENTS 30
#define PORT 8888 

void add_child_sockets_to_set(fd_set *readfds, int *client_sockets, int *max_sd);
void handle_clients_messages(int *client_sockets, fd_set *readfds, struct sockaddr_in *address);
void handle_new_client(int *new_socket, int *master_socket, int *client_sockets, struct sockaddr_in *address);
void setup_master_socket(int *master_socket, int *opt, struct sockaddr_in *address);
