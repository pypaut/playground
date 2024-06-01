#include <netinet/in.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>
#define PORT 8080

#include "server.h"

int main() {
    printf("Serving here\n");
    // Create server
    // Server.serve() or something
        // Receive each player data
        // Send other players' position to each player
    return 0;
}

server *create_server() {
    server *s = calloc(1, sizeof(server));

    s->players = calloc(1, sizeof(*player));
    s->nb_players = 0;

    return s;
}

void destroy_server(server *s) {
    for (size_t i = 0; i < s->nb_players; i++) {
        free(s->players[i]);
    }

    free(s->players);
    free(s);
}

int serve(server *s) {
    int server_fd, new_socket;
    ssize_t valread;
    struct sockaddr_in address;
    int opt = 1;
    socklen_t addrlen = sizeof(address);
    char buffer[1024] = { 0 };
 
    // Creating socket file descriptor
    if ((server_fd = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
        perror("socket failed");
        exit(EXIT_FAILURE);
    }
 
    // Set socket options
    if (setsockopt(server_fd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt))) {
        perror("setsockopt");
        exit(EXIT_FAILURE);
    }

    // Setup server address
    address.sin_family = AF_INET;
    address.sin_addr.s_addr = INADDR_ANY;
    address.sin_port = htons(PORT);
 
    // Bind socket to address
    if (bind(server_fd, (struct sockaddr*)&address, sizeof(address)) < 0) {
        perror("bind failed");
        exit(EXIT_FAILURE);
    }

    // Listen to client
    if (listen(server_fd, 3) < 0) {
        perror("listen");
        exit(EXIT_FAILURE);
    }

    // Connect to client socket
    if ((new_socket = accept(server_fd, (struct sockaddr*)&address, &addrlen)) < 0) {
        perror("accept");
        exit(EXIT_FAILURE);
    }

    read(new_socket, buffer, 1024 - 1); // subtract \0
    printf("%s\n", buffer);
 
    // Close the connected socket
    close(new_socket);

    // Close the listening socket
    close(server_fd);

    return 0;
}
