#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <errno.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <sys/time.h>

#include "server.h"

int main() { 
    int opt = 1; 
    int master_socket, new_socket, client_sockets[MAX_CLIENTS], activity, i;
    int max_sd; 
    struct sockaddr_in address; 

    // Set of socket descriptors
    fd_set readfds;

    // Initialise all client_sockets[] to 0 so not checked
    for (i = 0; i < MAX_CLIENTS; i++) {
        client_sockets[i] = 0;
    }

    // Master socket is the server's listening socket
    setup_master_socket(&master_socket, &opt, &address);
    puts("Waiting for connections ...");

    for (;;) {
        // Clear the socket set
        FD_ZERO(&readfds);

        // Add master socket to set
        FD_SET(master_socket, &readfds);
        max_sd = master_socket;

        // Add child sockets to set
        add_child_sockets_to_set(&readfds, client_sockets, &max_sd);

        // Wait for an activity on one of the sockets, timeout is NULL,
        // so wait indefinitely
        activity = select( max_sd + 1, &readfds, NULL, NULL, NULL);
        if ((activity < 0) && (errno!=EINTR)) {
            printf("select error");
        }

        if (FD_ISSET(master_socket, &readfds)) {
            handle_new_client(&new_socket, &master_socket, client_sockets, &address);
        } else {
            handle_clients_messages(client_sockets, &readfds, &address);
        }
    }

    return 0;
}

void setup_master_socket(int *master_socket, int *opt, struct sockaddr_in *address) {
    // Create a master socket
    if ((*master_socket = socket(AF_INET, SOCK_STREAM, 0)) == 0) {
        perror("socket failed");
        exit(EXIT_FAILURE);
    }

    // Set master socket to allow multiple connections,
    // this is just a good habit, it will work without this
    if (setsockopt(*master_socket, SOL_SOCKET, SO_REUSEADDR, (char*)opt, sizeof(opt)) < 0) {
        perror("setsockopt");
        exit(EXIT_FAILURE);
    }

    // Type of socket created
    address->sin_family = AF_INET;
    address->sin_addr.s_addr = INADDR_ANY;
    address->sin_port = htons(PORT);

    // Bind the socket to localhost port 8888
    if (bind(*master_socket, (struct sockaddr*)address, sizeof(*address)) < 0) {
        perror("bind failed");
        exit(EXIT_FAILURE);
    }
    printf("Listener on port %d\n", PORT);

    // Try to specify maximum of 3 pending connections for the master socket
    if (listen(*master_socket, 3) < 0) {
        perror("listen");
        exit(EXIT_FAILURE);
    }
}

void add_child_sockets_to_set(fd_set *readfds, int *client_sockets, int *max_sd) {
    for (size_t i = 0; i < MAX_CLIENTS; i++) {
        // Socket descriptor
        int sd = client_sockets[i];

        // If valid socket descriptor then add to read list
        if (sd > 0)
            FD_SET(sd, readfds);

        // Highest file descriptor number, need it for the select function
        if (sd > *max_sd)
            *max_sd = sd;
    }
}

void handle_new_client(int *new_socket, int *master_socket, int *client_sockets, struct sockaddr_in *address) {
    int addrlen = sizeof(*address);
    *new_socket = accept(*master_socket, (struct sockaddr*)address, (socklen_t*)&addrlen);
    if (*new_socket < 0) {
        perror("accept");
        exit(EXIT_FAILURE);
    }

    // Greetings message
    char *greetings = "greetings\0";
    size_t res = send(*new_socket, greetings, strlen(greetings), 0);
    if (res != strlen(greetings)) {
        perror("send");
    }

    // Add new socket to array of sockets
    for (int i = 0; i < MAX_CLIENTS; i++) {
        if (client_sockets[i] == 0) {
            client_sockets[i] = *new_socket;
            printf(
                "[ID: %d, IP: %s, PORT: %d]: connected\n",
                i, inet_ntoa(address->sin_addr), ntohs(address->sin_port)
            );
            break;
        }
    }
}

void handle_clients_messages(int *client_sockets, fd_set *readfds, struct sockaddr_in *address) {
    int addrlen = sizeof(address);
    char buffer[1024] = { 0 };
    char *message = "message\0";

    for (size_t i = 0; i < MAX_CLIENTS; i++) {
        int sd = client_sockets[i];

        if (FD_ISSET(sd, readfds)) {
            // If empty message, closing connection
            if (read(sd, buffer, 1024) == 0) {
                // Somebody disconnected , get his details and print
                getpeername(sd, (struct sockaddr*)&address, (socklen_t*)&addrlen);
                printf(
                    "[ID: %ld, IP: %s, PORT: %d]: disconnected\n",
                    i, inet_ntoa(address->sin_addr), ntohs(address->sin_port)
                );

                // Close the socket and mark as 0 in list for reuse
                close(sd);
                client_sockets[i] = 0;
            } else {
                // The buffer contains data from client
                printf(
                    "[ID: %ld, IP: %s, PORT: %d]: read \"%s\"\n",
                    i, inet_ntoa(address->sin_addr), ntohs(address->sin_port), buffer
                );
                send(sd, message, strlen(message), 0);
            }
        }
    }
}
